package v2022_04_15_preview

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/loadtestservice/2022-04-15-preview/loadtests"
)

type Client struct {
	LoadTests *loadtests.LoadTestsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	loadTestsClient := loadtests.NewLoadTestsClientWithBaseURI(endpoint)
	configureAuthFunc(&loadTestsClient.Client)

	return Client{
		LoadTests: &loadTestsClient,
	}
}
