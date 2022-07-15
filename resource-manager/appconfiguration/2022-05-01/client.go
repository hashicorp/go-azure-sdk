package v2022_05_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/appconfiguration/2022-05-01/configurationstores"
	"github.com/hashicorp/go-azure-sdk/resource-manager/appconfiguration/2022-05-01/deletedconfigurationstores"
	"github.com/hashicorp/go-azure-sdk/resource-manager/appconfiguration/2022-05-01/keyvalues"
	"github.com/hashicorp/go-azure-sdk/resource-manager/appconfiguration/2022-05-01/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/appconfiguration/2022-05-01/privatelinkresources"
)

type Client struct {
	ConfigurationStores        *configurationstores.ConfigurationStoresClient
	DeletedConfigurationStores *deletedconfigurationstores.DeletedConfigurationStoresClient
	KeyValues                  *keyvalues.KeyValuesClient
	PrivateEndpointConnections *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources       *privatelinkresources.PrivateLinkResourcesClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	configurationStoresClient := configurationstores.NewConfigurationStoresClientWithBaseURI(endpoint)
	configureAuthFunc(&configurationStoresClient.Client)

	deletedConfigurationStoresClient := deletedconfigurationstores.NewDeletedConfigurationStoresClientWithBaseURI(endpoint)
	configureAuthFunc(&deletedConfigurationStoresClient.Client)

	keyValuesClient := keyvalues.NewKeyValuesClientWithBaseURI(endpoint)
	configureAuthFunc(&keyValuesClient.Client)

	privateEndpointConnectionsClient := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&privateEndpointConnectionsClient.Client)

	privateLinkResourcesClient := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI(endpoint)
	configureAuthFunc(&privateLinkResourcesClient.Client)

	return Client{
		ConfigurationStores:        &configurationStoresClient,
		DeletedConfigurationStores: &deletedConfigurationStoresClient,
		KeyValues:                  &keyValuesClient,
		PrivateEndpointConnections: &privateEndpointConnectionsClient,
		PrivateLinkResources:       &privateLinkResourcesClient,
	}
}
