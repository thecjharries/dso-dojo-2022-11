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

.PHONY: test
test:
	npm test
	rm -rf cdktf.out
	cdktf synth
	go test
