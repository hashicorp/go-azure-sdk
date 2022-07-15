package v2016_06_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2016-06-01/connectiongateways"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2016-06-01/connections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2016-06-01/customapis"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2016-06-01/managedapis"
)

type Client struct {
	ConnectionGateways *connectiongateways.ConnectionGatewaysClient
	Connections        *connections.ConnectionsClient
	CustomAPIs         *customapis.CustomAPIsClient
	ManagedAPIs        *managedapis.ManagedAPIsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	connectionGatewaysClient := connectiongateways.NewConnectionGatewaysClientWithBaseURI(endpoint)
	configureAuthFunc(&connectionGatewaysClient.Client)

	connectionsClient := connections.NewConnectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&connectionsClient.Client)

	customAPIsClient := customapis.NewCustomAPIsClientWithBaseURI(endpoint)
	configureAuthFunc(&customAPIsClient.Client)

	managedAPIsClient := managedapis.NewManagedAPIsClientWithBaseURI(endpoint)
	configureAuthFunc(&managedAPIsClient.Client)

	return Client{
		ConnectionGateways: &connectionGatewaysClient,
		Connections:        &connectionsClient,
		CustomAPIs:         &customAPIsClient,
		ManagedAPIs:        &managedAPIsClient,
	}
}
