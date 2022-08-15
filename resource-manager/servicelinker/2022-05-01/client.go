package v2022_05_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicelinker/2022-05-01/links"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicelinker/2022-05-01/servicelinker"
)

type Client struct {
	Links         *links.LinksClient
	Servicelinker *servicelinker.ServicelinkerClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	linksClient := links.NewLinksClientWithBaseURI(endpoint)
	configureAuthFunc(&linksClient.Client)

	servicelinkerClient := servicelinker.NewServicelinkerClientWithBaseURI(endpoint)
	configureAuthFunc(&servicelinkerClient.Client)

	return Client{
		Links:         &linksClient,
		Servicelinker: &servicelinkerClient,
	}
}
