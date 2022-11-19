
import { Construct } from "constructs";
import { App, TerraformStack, TerraformOutput } from "cdktf";
import { AwsProvider } from "@cdktf/provider-aws/lib/provider";
import { Instance } from "@cdktf/provider-aws/lib/instance";
import { KeyPair } from "@cdktf/provider-aws/lib/key-pair";
import { TlsProvider } from "@cdktf/provider-tls/lib/provider";
import { PrivateKey } from "@cdktf/provider-tls/lib/private-key";
import { DataAwsAmi } from "@cdktf/provider-aws/lib/data-aws-ami";
import { DataAwsSecurityGroup } from "@cdktf/provider-aws/lib/data-aws-security-group";

export class DsoDojo202211 extends TerraformStack {
    constructor(scope: Construct, id: string) {
        super(scope, id);

        new AwsProvider(this, "aws", {
            region: "us-west-2",
        });

        const tlsProvider = new TlsProvider(this, "null", {});

        // Don't do this in production
        // https://registry.terraform.io/providers/hashicorp/tls/latest/docs/resources/private_key
        const key = new PrivateKey(this, "key", {
            algorithm: "RSA",
            rsaBits: 4096,
            provider: tlsProvider,
        });

        const keyPair = new KeyPair(this, "keypair", {
            keyName: "DSO Dojo 2022-11",
            publicKey: key.publicKeyOpenssh,
        });

        // Don't do this in production
        // https://registry.terraform.io/providers/hashicorp/tls/latest/docs/resources/private_key
        new TerraformOutput(this, "public_key", {
            value: key.publicKeyPem,
        });

        // Don't do this in production
        // https://registry.terraform.io/providers/hashicorp/tls/latest/docs/resources/private_key
        new TerraformOutput(this, "private_key", {
            value: key.privateKeyPem,
            sensitive: true,
        });

        const ami = new DataAwsAmi(this, "ami", {
            mostRecent: true,
            owners: ["self"],
            filter: [
                {
                    name: "name",
                    values: ["dso-dojo-202211"],
                },
            ],
        });

        new TerraformOutput(this, "ami_id", {
            value: ami.id,
        });

        const securityGroup = new DataAwsSecurityGroup(this, "security_group", {
            name: "default",
        });

        new TerraformOutput(this, "security_group_id", {
            value: securityGroup.id,
        });

        const ec2Instance = new Instance(this, "compute", {
            ami: ami.id,
            instanceType: "t2.micro",
            keyName: keyPair.keyName,
            associatePublicIpAddress: true,
            securityGroups: [securityGroup.name],
        });

        new TerraformOutput(this, "ec2_id", {
            value: ec2Instance.id,
        });

        new TerraformOutput(this, "ec2_public_ip", {
            value: ec2Instance.publicIp,
        });
    }
}

const app = new App();
new DsoDojo202211(app, "terraform");
app.synth();
