package middleware

import (
	"net"
	"net/http"

	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/svbh/middleware/apiserver/cmd/apiserver/gen"
	"github.com/svbh/middleware/types"
)

// RestServer is the REST api server
type RestServer struct {
	listenURL           string
	listener            net.Listener
	httpServer          *http.Server
	requirementResolver types.RequirementResolver
	idToTransactorLUT   map[uint64]*bind.TransactOpts
	donor, vendor       *bind.TransactOpts
}
type routeAddFunc func(*mux.Router, *RestServer)

func NewAPIGateway(listenURL string) (*RestServer, error) {
	// create server instance
	srv := RestServer{
		listenURL: listenURL,
	}
	alloc := make(core.GenesisAlloc)
	srv.idToTransactorLUT = make(map[uint64]*bind.TransactOpts)

	dKey, _ := crypto.GenerateKey()
	donor := bind.NewKeyedTransactor(dKey)

	vKey, _ := crypto.GenerateKey()
	vendor := bind.NewKeyedTransactor(vKey)

	for _, n := range mockListings {
		key, _ := crypto.GenerateKey()
		ngo := bind.NewKeyedTransactor(key)
		srv.idToTransactorLUT[n.ID] = ngo
		alloc[ngo.From] = core.GenesisAccount{Balance: big.NewInt(133700000)}
		//opts := &bind.TransactOpts{From: ngo.From, Signer: ngo.Signer}
		//_, err = nonProfit.CreateRequest(opts, vendor.From, ngo.From)
		//
		//
		//callopts := &bind.CallOpts{
		//	Pending: true,
		//}
	}
	srv.vendor = vendor
	srv.donor = donor
	alloc[donor.From] = core.GenesisAccount{Balance: big.NewInt(133700000)}
	alloc[vendor.From] = core.GenesisAccount{Balance: big.NewInt(133700000)}

	cKey, _ := crypto.GenerateKey()
	ctxAccount := bind.NewKeyedTransactor(cKey)
	alloc[ctxAccount.From] = core.GenesisAccount{Balance: big.NewInt(133700000)}

	sim := backends.NewSimulatedBackend(alloc)

	_, _, _, err := gen.DeployNonProfit(ctxAccount, sim)
	if err != nil {
		log.Fatalf("Failed to deploy new token contract: %v", err)
	}

	//ids,vendorIDs,donorIDs,statuses,err := nonProfit.GetRequestsByNonProfitId(callopts, ngo.From)
	//
	//log.Infof("ID: %v", ids)
	//log.Infof("VENDOR ID: %v", vendorIDs)
	//log.Infof("DONOR ID: %v", donorIDs)
	//log.Infof("Statuses: %v", statuses)
	//log.Infof("ERROR: %v", err)
	sim.Commit()

	// if no URL was specified, just return (used during unit/integ tests)
	if listenURL == "" {
		return &srv, nil
	}

	// setup the top level routes
	router := mux.NewRouter()
	prefixRoutes := map[string]routeAddFunc{
		"/api/listings/": addAPIListingsRoute,
	}

	for prefix, subRouter := range prefixRoutes {
		sub := router.PathPrefix(prefix).Subrouter().StrictSlash(true)
		subRouter(sub, &srv)
	}

	log.Infof("Starting server at %s", listenURL)

	// listener
	listener, err := net.Listen("tcp", listenURL)
	if err != nil {
		log.Errorf("Error starting listener. Err: %v", err)
		return nil, err
	}
	srv.listener = listener

	// create a http server
	srv.httpServer = &http.Server{Addr: listenURL, Handler: router}
	go srv.httpServer.Serve(listener)

	return &srv, nil
}

// Stop stops the http server
func (s *RestServer) Stop() error {
	if s.httpServer != nil {
		s.httpServer.Close()
	}

	return nil
}
