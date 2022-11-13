# DSO Dojo 2022-11

## Overview

One of the projects I've been running at work is a DevSecOps and SRE Dojo. This repo is my solution for our November 2022 exercise.

### Exercise

The stated goal was to get [Gin's Quick Start](https://github.com/gin-gonic/gin#quick-start) on an EC2 server.

### Toolchain

My initial goal is to do the following

1. Create a build pipeline in GitHub Actions that will create a GitHub Release on tag events.
2. Use CDKTF to provision LocalStack.
3. Terratest everything for target state.

Right now I'm not sure about provisioning the the EC2 instance. Keeping with "do the most complicated thing" approach, Packer to create an AMI sounds pretty neat.

## Usage

### Prerequisites

You'll need a LocalStack Pro (trial) account. You can get one [here](https://localstack.cloud/).

## Notes

### Rootless Podman

... is a really bad idea. LocalStack [doesn't exactly support Podman](https://docs.localstack.cloud/localstack/podman/) and requires a ton of fiddling to make happy if you're following a standard Arch rootless setup.

* You really need [an LS profile](https://docs.localstack.cloud/localstack/configuration/#profiles) and
* you need to understand [the Docker flags](https://docs.localstack.cloud/localstack/configuration/#docker)
* because you'll probably need to add [this important Podman flag](https://github.com/containers/podman/issues/14284#issuecomment-1130113553).

I gave up and installed Docker on Arch because I spent three hours pulling out my hair. [This issue is the root cause](https://github.com/containers/podman/issues/4900) and seems to have been erroneously closed (or they don't have good regression testing). I did all sorts of things and just couldn't get it to work because of VirtualBox USB `/dev`ices.

### LocalStack Pro Required for Packer

Apparently actually [doing AMI things is behind LocalStack Pro](https://github.com/localstack/localstack/issues/1996). Packer does seem to work pretty well minus that part provided you add this line to your `builder`:

```json
"custom_endpoint_ec2": "http://localhost:4566"
```

~~Since one of my goals was to do this without creating any accounts I'll move on for now.~~

### Lots to learn with CDKTF

Following the CDKTF quick start sets up a stack. Apparently these don't work with [resource assertions](https://github.com/hashicorp/terraform-cdk/issues/1850#issuecomment-1153883827).
