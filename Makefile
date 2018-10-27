# Usage:
# $ make == $ make build
# $ make install

.PHONY: dep client install 

.DEFAULT_GOAL := build
ETH_CONTRACTS := onchain/eth/contracts
GENERATED_FILES := $(filter-out $(shell find $(ETH_CONTRACTS) -name '*_test.go'), $(shell find $(ETH_CONTRACTS) -name '*.go'))


build: dep client


dep:
	dep ensure -vendor-only


client: gen
	go build -o client


gen:
	abigen -sol $(ETH_CONTRACTS)/DOSAddressBridge.sol --pkg dosproxy --out $(ETH_CONTRACTS)/DOSAddressBridge.go
	abigen -sol $(ETH_CONTRACTS)/DOSOnChainSDK.sol --pkg dosproxy --out $(ETH_CONTRACTS)/DOSOnChainSDK.go
	abigen -sol $(ETH_CONTRACTS)/DOSProxy.sol --pkg dosproxy --out $(ETH_CONTRACTS)/DOSProxy.go
	cp $(ETH_CONTRACTS)/DOSOnChainSDK.sol $(ETH_CONTRACTS)/userContract/
	cp $(ETH_CONTRACTS)/Ownable.sol $(ETH_CONTRACTS)/userContract/
	abigen -sol $(ETH_CONTRACTS)/userContract/AskMeAnything.sol --pkg userContract --out $(ETH_CONTRACTS)/userContract/AskMeAnything.go
	rm $(ETH_CONTRACTS)/userContract/DOSOnChainSDK.sol
	rm $(ETH_CONTRACTS)/userContract/Ownable.sol


install: dep client
	go install


deploy: gen
	go run onchain/eth/deploy.go -contractPath $(ETH_CONTRACTS) -step ProxyAndBridg
	go run onchain/eth/deploy.go -contractPath $(ETH_CONTRACTS) -step SDKAndAMA
	go run onchain/eth/deploy.go -contractPath $(ETH_CONTRACTS) -step SetProxyAddress


clean:
	rm -f client
	rm -f $(GENERATED_FILES)
