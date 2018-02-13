.PHONY: contracts protocol

contracts:
	npm install
	mkdir -p abi
	mkdir -p contracts
	$(call GenerateBinding,BasicElection)
	$(call GenerateBinding,TieredElection)
	$(call GenerateBinding,TieredBallot)
	$(call GenerateBinding,BasePool)
	$(call GenerateBinding,BaseBallot)
	$(call GenerateBinding,BaseElection)
	rm -rf abi

define GenerateBinding
	cat "./node_modules/@netvote/elections-solidity/build/contracts/$(1).json" |jq '.abi' > "abi/$(1).abi"
    abigen --abi "./abi/$(1).abi" --type "$(1)" --pkg contracts --out "contracts/$(1).go"
endef

build:
	go build cmd/tally/main.go

protocol:
	protoc --go_out=:./protocol -I. -I./vendor/github.com/netvote/elections-solidity/protocol ./vendor/github.com/netvote/elections-solidity/protocol/*.proto
	mv protocol/vendor/github.com/netvote/elections-solidity/protocol/*.go protocol/
	rm -rf protocol/vendor
	sed -i.bak 's/package netvote/package protocol/g' protocol/*.go
	rm protocol/*.bak
