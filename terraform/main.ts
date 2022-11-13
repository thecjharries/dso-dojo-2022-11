
import { Construct } from "constructs";
import { App, TerraformStack, TerraformOutput } from "cdktf";
import { AwsProvider } from "@cdktf/provider-aws/lib/provider";
import { Instance } from "@cdktf/provider-aws/lib/instance";
import { TlsProvider } from "@cdktf/provider-tls/lib/provider";
import { PrivateKey } from "@cdktf/provider-tls/lib/private-key";

export class MyStack extends TerraformStack {
    constructor(scope: Construct, id: string) {
        super(scope, id);

        new AwsProvider(this, "aws", {
            region: "us-west-2",
            secretKey: "test",
            accessKey: "test",
            skipCredentialsValidation: true,
            skipMetadataApiCheck: "true",
            skipRequestingAccountId: true,
            s3UsePathStyle: true,
            endpoints: [
                {
                    ec2: "http://localhost:4566",
                }
            ]
        });

        const tlsProvider = new TlsProvider(this, "null", {});

        const key = new PrivateKey(this, "key", {
            algorithm: "RSA",
            rsaBits: 4096,
            provider: tlsProvider,
        });

        const ec2Instance = new Instance(this, "compute", {
            ami: "ami-01456a894f71116f2",
            instanceType: "t2.micro",
        });

        new TerraformOutput(this, "public_ip", {
            value: ec2Instance.publicIp,
        });
    }
}

const app = new App();
new MyStack(app, "terraform");
app.synth();
