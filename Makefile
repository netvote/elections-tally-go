.PHONY: contracts

contracts:
	npm install
	mkdir -p abi
	mkdir -p contracts
	$(call GenerateBinding,BasicElection)
	$(call GenerateBinding,TieredElection)
	$(call GenerateBinding,TieredBallot)
	$(call GenerateBinding,TieredPool)
	rm -rf abi

define GenerateBinding
	cat "./node_modules/@netvote/elections-solidity/build/contracts/$(1).json" |jq '.abi' > "abi/$(1).abi"
    abigen --abi "./abi/$(1).abi" --type "$(1)" --pkg contracts --out "contracts/$(1).go"
endef

build:
	go build cmd/tally/main.go
