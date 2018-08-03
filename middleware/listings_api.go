package middleware

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/svbh/middleware/types"
)

var mockListings = []types.Request{
	{
		ID:            1,
		Description:   "Support Immigrant children at the border",
		Summary:       "This will help with one week's worth of baby formula",
		NonProfitName: "RAICES, Texas",
		Status:        types.OPEN,
		Price:         "2 ETH",
		Category:      []string{"immigration", "children", "human rights"},
	},
	{
		ID:            2,
		Description:   "Unite separated immigration children with their parents",
		Summary:       "This will pay for the lawyer fees for one family",
		NonProfitName: "American Civil Liberties Union",
		Status:        types.OPEN,
		Price:         "30 ETH",
		Category:      []string{"immigration", "children", "human rights"},
	},
	{
		ID:            3,
		Description:   "School supplies",
		Summary:       "This will sponsor one child's academic year worth of school supplies",
		NonProfitName: "Teach for America",
		Status:        types.OPEN,
		Price:         "25 ETH",
		Category:      []string{"education", "children"},
	},
	{
		ID:            4,
		Description:   "Tech education for women",
		Summary:       "We provide 1 year's subscription for Udacity for 5 girls from underrepresented communities",
		NonProfitName: "Women Who code",
		Status:        types.OPEN,
		Price:         "42 ETH",
		Category:      []string{"education", "diversity", "tech"},
	},
}

// addAPIListingsRoute adds Interface routes
func addAPIListingsRoute(r *mux.Router, srv *RestServer) {
	r.Methods("GET").Subrouter().HandleFunc("/", MakeHTTPHandler(srv.listInterfaceHandler))
}

func (s *RestServer) listInterfaceHandler(r *http.Request) (interface{}, error) {
	return makeMockData()
}
func makeMockData() ([]types.Request, error) {
	return mockListings, nil
}
