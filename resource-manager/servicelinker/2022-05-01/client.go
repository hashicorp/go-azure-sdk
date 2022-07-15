package v2022_05_01

import (
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicelinker/2022-05-01/links"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicelinker/2022-05-01/servicelinker"
)

type Client struct {
	Links         *links.LinksClient
	ServiceLinker *servicelinker.ServiceLinkerClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	linksClient := links.NewLinksClientWithBaseURI(endpoint)
	configureAuthFunc(&linksClient.Client)

	serviceLinkerClient := servicelinker.NewServiceLinkerClientWithBaseURI(endpoint)
	configureAuthFunc(&serviceLinkerClient.Client)

	return Client{
		Links:         &linksClient,
		ServiceLinker: &serviceLinkerClient,
	}
}
