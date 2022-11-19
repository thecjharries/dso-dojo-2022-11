package test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/retry"
	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"

	"github.com/stretchr/testify/assert"

	"golang.org/x/crypto/ssh"
)

type PingResponse struct {
	Message string `json:"message"`
}

func TestDsoDojo202211(t *testing.T) {
	tempTestFolder := test_structure.CopyTerraformFolderToTemp(t, "./cdktf.out/stacks/terraform", ".")
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: tempTestFolder,
	})

	terraform.InitAndApply(t, terraformOptions)

	defer terraform.Destroy(t, terraformOptions)

	ec2Ip := terraform.Output(t, terraformOptions, "ec2_public_ip")
	assert.NotEmpty(t, ec2Ip, "An instance should have been created")

	privateKey := terraform.Output(t, terraformOptions, "private_key")
	assert.NotEmpty(t, privateKey, "A private key should exist")

	host := ec2Ip + ":22"
	user := "ubuntu"
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

	_, err = retry.DoWithRetryE(t, "Running SSH test", 6, 10*time.Second, func() (string, error) {
		client, err := ssh.Dial("tcp", host, conf)
		if err != nil {
			return "", err
		}
		defer client.Close()

		session, err := client.NewSession()
		assert.NoError(t, err, "Failed to create session")
		defer session.Close()

		hostname, err := session.Output("hostname")
		assert.NoError(t, err, "Failed to get hostname")
		logger.Log(t, "Remote hostname: ", string(hostname))
		assert.NotEmpty(t, string(hostname), "A hostname should exist")
		return "", nil
	})
	assert.NoError(t, err, "Failed to SSH")

	_, err = retry.DoWithRetryE(t, "Running GET test", 6, 10*time.Second, func() (string, error) {
		reponse, err := http.Get("http://" + ec2Ip + "/ping")
		if err != nil {
			return "", err
		}
		defer reponse.Body.Close()

		body, err := ioutil.ReadAll(reponse.Body)
		assert.NoError(t, err, "Failed to read response")
		var pingResponse PingResponse
		err = json.Unmarshal(body, &pingResponse)
		assert.NoError(t, err, "Failed to unmarshal response")
		assert.Equal(t, "pong", pingResponse.Message, "Ping response should be pong")
		return "", nil
	})
	assert.NoError(t, err, "Failed to hit server")
}
