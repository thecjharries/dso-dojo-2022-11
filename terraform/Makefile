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

.PHONY: ssh
ssh: ssh_key
	ssh -i ssh_key -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null -p $(shell docker ps -a --format '{{ json . }}' | jq -s . | jq '.[] | select(.Names | contains("'$(shell jq -r '.terraform.ec2_id' outputs.json)'")) | .Ports' | sed -E 's@^.*0.0:(.*)->22/tcp,.*@\1@') root@127.0.0.1

.PHONY: test
test:
	npm test
	rm -rf cdktf.out
	cdktf synth
	go test
