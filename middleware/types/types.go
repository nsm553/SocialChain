package types

import "net/http"

const (
	OPEN = iota
	CLOSED
)

type ListingStatus int

// HTTPAPIFunc Type for HTTP handler functions
type HTTPAPIFunc func(r *http.Request) (interface{}, error)

type Contract struct {
	ListingID        uint64
	VendorAddress    string
	DonorAddress     string
	NonProfitAddress string
	Status           bool
}

type RequirementResolver interface {
	CreateRequirement(c *Contract) error
	UpdateRequirement() error
	SetVendorBid(bid int) error
	SetDonor(donorAddress string) error
}

type Request struct {
	ID            uint64        `json:"id"`
	Description   string        `json:"description"`
	Summary       string        `json:"summary"`
	NonProfitName string        `json:"name"`
	Status        ListingStatus `json:"listing-status"`
	Price         string        `json:"price"`
	Category      []string      `json:"category"`
}
