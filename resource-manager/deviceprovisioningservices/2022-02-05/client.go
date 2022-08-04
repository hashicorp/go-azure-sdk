package v2022_02_05

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/deviceprovisioningservices/2022-02-05/delete"
	"github.com/hashicorp/go-azure-sdk/resource-manager/deviceprovisioningservices/2022-02-05/get"
	"github.com/hashicorp/go-azure-sdk/resource-manager/deviceprovisioningservices/2022-02-05/patch"
	"github.com/hashicorp/go-azure-sdk/resource-manager/deviceprovisioningservices/2022-02-05/post"
	"github.com/hashicorp/go-azure-sdk/resource-manager/deviceprovisioningservices/2022-02-05/put"
)

type Client struct {
	DELETE *delete.DELETEClient
	GET    *get.GETClient
	PATCH  *patch.PATCHClient
	POST   *post.POSTClient
	PUT    *put.PUTClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	dELETEClient := delete.NewDELETEClientWithBaseURI(endpoint)
	configureAuthFunc(&dELETEClient.Client)

	gETClient := get.NewGETClientWithBaseURI(endpoint)
	configureAuthFunc(&gETClient.Client)

	pATCHClient := patch.NewPATCHClientWithBaseURI(endpoint)
	configureAuthFunc(&pATCHClient.Client)

	pOSTClient := post.NewPOSTClientWithBaseURI(endpoint)
	configureAuthFunc(&pOSTClient.Client)

	pUTClient := put.NewPUTClientWithBaseURI(endpoint)
	configureAuthFunc(&pUTClient.Client)

	return Client{
		DELETE: &dELETEClient,
		GET:    &gETClient,
		PATCH:  &pATCHClient,
		POST:   &pOSTClient,
		PUT:    &pUTClient,
	}
}
