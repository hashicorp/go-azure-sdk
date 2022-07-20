package v2020_08_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/search/2020-08-01/adminkeys"
	"github.com/hashicorp/go-azure-sdk/resource-manager/search/2020-08-01/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/search/2020-08-01/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/search/2020-08-01/querykeys"
	"github.com/hashicorp/go-azure-sdk/resource-manager/search/2020-08-01/services"
	"github.com/hashicorp/go-azure-sdk/resource-manager/search/2020-08-01/sharedprivatelinkresources"
)

type Client struct {
	AdminKeys                  *adminkeys.AdminKeysClient
	PrivateEndpointConnections *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources       *privatelinkresources.PrivateLinkResourcesClient
	QueryKeys                  *querykeys.QueryKeysClient
	Services                   *services.ServicesClient
	SharedPrivateLinkResources *sharedprivatelinkresources.SharedPrivateLinkResourcesClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	adminKeysClient := adminkeys.NewAdminKeysClientWithBaseURI(endpoint)
	configureAuthFunc(&adminKeysClient.Client)

	privateEndpointConnectionsClient := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&privateEndpointConnectionsClient.Client)

	privateLinkResourcesClient := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI(endpoint)
	configureAuthFunc(&privateLinkResourcesClient.Client)

	queryKeysClient := querykeys.NewQueryKeysClientWithBaseURI(endpoint)
	configureAuthFunc(&queryKeysClient.Client)

	servicesClient := services.NewServicesClientWithBaseURI(endpoint)
	configureAuthFunc(&servicesClient.Client)

	sharedPrivateLinkResourcesClient := sharedprivatelinkresources.NewSharedPrivateLinkResourcesClientWithBaseURI(endpoint)
	configureAuthFunc(&sharedPrivateLinkResourcesClient.Client)

	return Client{
		AdminKeys:                  &adminKeysClient,
		PrivateEndpointConnections: &privateEndpointConnectionsClient,
		PrivateLinkResources:       &privateLinkResourcesClient,
		QueryKeys:                  &queryKeysClient,
		Services:                   &servicesClient,
		SharedPrivateLinkResources: &sharedPrivateLinkResourcesClient,
	}
}
