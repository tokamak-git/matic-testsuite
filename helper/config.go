package helper

import (
	"log"

	"github.com/maticnetwork/bor/rpc"

	"github.com/maticnetwork/bor/ethclient"
)

//
// Get main/matic clients
//

// GetMainClient returns main chain's eth client
func GetMainClient() *ethclient.Client {
	// TODO: pass url from config
	mainRPCClient, err := rpc.Dial("https://goerli.infura.io/v3/a0358f426e6243528a53a8c47244c1a7")
	if err != nil {
		log.Fatalln("Unable to dial via ethClient", "URL=", "https://goerli.infura.io/v3/a0358f426e6243528a53a8c47244c1a7", "chain=eth", "Error", err)
	}

	mainChainClient := ethclient.NewClient(mainRPCClient)
	return mainChainClient
}

// GetMaticClient returns matic's eth client
func GetMaticClient() *ethclient.Client {
	// TODO: pass url from config
	maticRPCClient, err := rpc.Dial("https://rpc-mumbai.matic.today")
	if err != nil {
		log.Fatal(err)
	}

	maticClient := ethclient.NewClient(maticRPCClient)

	return maticClient
}
