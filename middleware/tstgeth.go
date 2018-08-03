package middleware

import (
	"github.com/melonproject/ethereum-go-client/client"
)

var ethAddress = "localhost:8545"

func ExampleClient() {
	ethclient := client.NewClient(ethAddress)

	type Inp struct	{
		detail 	int
		creator	string
		bids	int
	}
	// type := `{"detail": 123, "creator": "testGO", "bids": 234}`
	var input = json.Unmarshal([]byte &Inp)

	peerCount, err := ethclient.Endpoints.createRequirement(input)
	if err != nil {
		panic(err)
	}

	// peers, err := ethclient.Endpoints.Parity.NetPeers()
	// if err != nil {
	// 	panic(err)
	// }

	// if peerCount != peers.Active {
	// 	panic("Number of peers mismatch")
	// }
}
