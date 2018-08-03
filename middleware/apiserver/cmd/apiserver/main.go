package main

import (
	"github.com/svbh/middleware"
)

func main() {
	waitCh := make(chan bool)

	//dKey, _ := crypto.GenerateKey()
	//donor := bind.NewKeyedTransactor(dKey)
	//
	//vKey, _ := crypto.GenerateKey()
	//vendor := bind.NewKeyedTransactor(vKey)
	//
	//nKey, _ := crypto.GenerateKey()
	//ngo := bind.NewKeyedTransactor(nKey)
	//
	//
	//alloc := make(core.GenesisAlloc)
	//alloc[donor.From] = core.GenesisAccount{Balance: big.NewInt(133700000)}
	//alloc[ngo.From] = core.GenesisAccount{Balance: big.NewInt(133700000)}
	//alloc[vendor.From] = core.GenesisAccount{Balance: big.NewInt(133700000)}
	//
	//sim := backends.NewSimulatedBackend(alloc)
	//
	//_, _, nonProfit, err := gen.DeployNonProfit(ngo, sim)
	//if err != nil {
	//	log.Fatalf("Failed to deploy new token contract: %v", err)
	//}
	//
	//
	//opts := &bind.TransactOpts{From: ngo.From, Signer: ngo.Signer}
	//_, err = nonProfit.CreateRequest(opts, vendor.From, ngo.From)
	//
	//
	//callopts := &bind.CallOpts{
	//	Pending: true,
	//}
	//ids,vendorIDs,donorIDs,statuses,err := nonProfit.GetRequestsByNonProfitId(callopts, ngo.From)
	//
	//log.Infof("ID: %v", ids)
	//log.Infof("VENDOR ID: %v", vendorIDs)
	//log.Infof("DONOR ID: %v", donorIDs)
	//log.Infof("Statuses: %v", statuses)
	//log.Infof("ERROR: %v", err)
	//sim.Commit()

	middleware.NewAPIGateway("localhost:9007")
	// wait forever
	<-waitCh

}
