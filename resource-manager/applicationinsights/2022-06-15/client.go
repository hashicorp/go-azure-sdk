package v2022_06_15

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/applicationinsights/2022-06-15/webtestsapis"
)

type Client struct {
	WebTestsAPIs *webtestsapis.WebTestsAPIsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	webTestsAPIsClient := webtestsapis.NewWebTestsAPIsClientWithBaseURI(endpoint)
	configureAuthFunc(&webTestsAPIsClient.Client)

	return Client{
		WebTestsAPIs: &webTestsAPIsClient,
	}
}
