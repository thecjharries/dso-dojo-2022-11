.PHONY: prep
prep:
	docker pull ubuntu:focal
	docker tag ubuntu:focal localstack-ec2/ubuntu-focal-ami:ami-000001

.PHONY: deploy
deploy:
	cdktf deploy

outputs.json:
	cdktf outputs --outputs-file-include-sensitive-outputs --outputs-file outputs.json

ssh_key: outputs.json
	jq -r '.terraform.private_key' outputs.json > ssh_key
	chmod 0600 ssh_key

.PHONY: clean
clean:
	cdktf destroy
	docker rmi localstack-ec2/ubuntu-focal-ami:ami-000001
