package v2022_01_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/redisenterprise/2022-01-01/databases"
	"github.com/hashicorp/go-azure-sdk/resource-manager/redisenterprise/2022-01-01/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/redisenterprise/2022-01-01/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/redisenterprise/2022-01-01/redisenterprise"
)

type Client struct {
	Databases                  *databases.DatabasesClient
	PrivateEndpointConnections *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources       *privatelinkresources.PrivateLinkResourcesClient
	RedisEnterprise            *redisenterprise.RedisEnterpriseClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	databasesClient := databases.NewDatabasesClientWithBaseURI(endpoint)
	configureAuthFunc(&databasesClient.Client)

	privateEndpointConnectionsClient := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&privateEndpointConnectionsClient.Client)

	privateLinkResourcesClient := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI(endpoint)
	configureAuthFunc(&privateLinkResourcesClient.Client)

	redisEnterpriseClient := redisenterprise.NewRedisEnterpriseClientWithBaseURI(endpoint)
	configureAuthFunc(&redisEnterpriseClient.Client)

	return Client{
		Databases:                  &databasesClient,
		PrivateEndpointConnections: &privateEndpointConnectionsClient,
		PrivateLinkResources:       &privateLinkResourcesClient,
		RedisEnterprise:            &redisEnterpriseClient,
	}
}
