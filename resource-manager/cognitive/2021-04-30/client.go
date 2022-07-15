package v2021_04_30

import (
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2021-04-30/cognitiveservicesaccounts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2021-04-30/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2021-04-30/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2021-04-30/skus"
)

type Client struct {
	CognitiveServicesAccounts  *cognitiveservicesaccounts.CognitiveServicesAccountsClient
	PrivateEndpointConnections *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources       *privatelinkresources.PrivateLinkResourcesClient
	Skus                       *skus.SkusClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	cognitiveServicesAccountsClient := cognitiveservicesaccounts.NewCognitiveServicesAccountsClientWithBaseURI(endpoint)
	configureAuthFunc(&cognitiveServicesAccountsClient.Client)

	privateEndpointConnectionsClient := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&privateEndpointConnectionsClient.Client)

	privateLinkResourcesClient := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI(endpoint)
	configureAuthFunc(&privateLinkResourcesClient.Client)

	skusClient := skus.NewSkusClientWithBaseURI(endpoint)
	configureAuthFunc(&skusClient.Client)

	return Client{
		CognitiveServicesAccounts:  &cognitiveServicesAccountsClient,
		PrivateEndpointConnections: &privateEndpointConnectionsClient,
		PrivateLinkResources:       &privateLinkResourcesClient,
		Skus:                       &skusClient,
	}
}
