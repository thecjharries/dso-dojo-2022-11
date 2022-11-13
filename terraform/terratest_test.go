package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"

	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"

	"github.com/stretchr/testify/assert"

	"golang.org/x/crypto/ssh"
)

func TestTerraformMyStack(t *testing.T) {
	tempTestFolder := test_structure.CopyTerraformFolderToTemp(t, "./cdktf.out/stacks/terraform", ".")
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: tempTestFolder,
	})

	terraform.InitAndApply(t, terraformOptions)

	ec2Id := terraform.Output(t, terraformOptions, "ec2_id")
	assert.NotEmpty(t, ec2Id, "An instance should have been created")

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	assert.NoError(t, err, "Failed to create docker client")
	defer cli.Close()

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	assert.NoError(t, err, "Failed to list containers")

	var sshPort uint16

	for _, container := range containers {
		for _, name := range container.Names {
			if "/localstack-ec2."+ec2Id == name {
				for _, port := range container.Ports {
					if port.PrivatePort == 22 {
						sshPort = port.PublicPort
						break
					}
				}
			}
			if 0 != sshPort {
				break
			}
		}
		if 0 != sshPort {
			break
		}
	}
	logger.Log(t, "sshPort: ", sshPort)

	privateKey := terraform.Output(t, terraformOptions, "private_key")
	assert.NotEmpty(t, privateKey, "A private key should exist")

	host := "127.0.0.1:" + fmt.Sprint(sshPort)
	user := "root"
	pKey := []byte(privateKey)

	signer, err := ssh.ParsePrivateKey(pKey)
	assert.NoError(t, err, "Failed to parse private key")

	conf := &ssh.ClientConfig{
		User:            user,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
	}

	conn, err := ssh.Dial("tcp", host, conf)
	assert.NoError(t, err, "Failed to dial")
	defer conn.Close()

	session, err := conn.NewSession()
	assert.NoError(t, err, "Failed to create session")
	defer session.Close()

	hostname, err := session.Output("hostname")
	assert.NoError(t, err, "Failed to get hostname")
	logger.Log(t, "Remote hostname: ", string(hostname))
	assert.NotEmpty(t, string(hostname), "A hostname should exist")
}
