package v2022_05_13

import (
	"github.com/hashicorp/go-azure-sdk/resource-manager/confidentialledger/2022-05-13/confidentialledger"
	"github.com/hashicorp/go-azure-sdk/resource-manager/confidentialledger/2022-05-13/nameavailability"
)

type Client struct {
	ConfidentialLedger *confidentialledger.ConfidentialLedgerClient
	NameAvailability   *nameavailability.NameAvailabilityClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	confidentialLedgerClient := confidentialledger.NewConfidentialLedgerClientWithBaseURI(endpoint)
	configureAuthFunc(&confidentialLedgerClient.Client)

	nameAvailabilityClient := nameavailability.NewNameAvailabilityClientWithBaseURI(endpoint)
	configureAuthFunc(&nameAvailabilityClient.Client)

	return Client{
		ConfidentialLedger: &confidentialLedgerClient,
		NameAvailability:   &nameAvailabilityClient,
	}
}
