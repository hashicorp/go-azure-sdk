package v2023_01_31

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/digitaltwins/2023-01-31/checknameavailability"
	"github.com/hashicorp/go-azure-sdk/resource-manager/digitaltwins/2023-01-31/digitaltwinsinstance"
	"github.com/hashicorp/go-azure-sdk/resource-manager/digitaltwins/2023-01-31/endpoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/digitaltwins/2023-01-31/privateendpoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/digitaltwins/2023-01-31/timeseriesdatabaseconnections"
)

type Client struct {
	CheckNameAvailability         *checknameavailability.CheckNameAvailabilityClient
	DigitalTwinsInstance          *digitaltwinsinstance.DigitalTwinsInstanceClient
	Endpoints                     *endpoints.EndpointsClient
	PrivateEndpoints              *privateendpoints.PrivateEndpointsClient
	TimeSeriesDatabaseConnections *timeseriesdatabaseconnections.TimeSeriesDatabaseConnectionsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	checkNameAvailabilityClient := checknameavailability.NewCheckNameAvailabilityClientWithBaseURI(endpoint)
	configureAuthFunc(&checkNameAvailabilityClient.Client)

	digitalTwinsInstanceClient := digitaltwinsinstance.NewDigitalTwinsInstanceClientWithBaseURI(endpoint)
	configureAuthFunc(&digitalTwinsInstanceClient.Client)

	endpointsClient := endpoints.NewEndpointsClientWithBaseURI(endpoint)
	configureAuthFunc(&endpointsClient.Client)

	privateEndpointsClient := privateendpoints.NewPrivateEndpointsClientWithBaseURI(endpoint)
	configureAuthFunc(&privateEndpointsClient.Client)

	timeSeriesDatabaseConnectionsClient := timeseriesdatabaseconnections.NewTimeSeriesDatabaseConnectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&timeSeriesDatabaseConnectionsClient.Client)

	return Client{
		CheckNameAvailability:         &checkNameAvailabilityClient,
		DigitalTwinsInstance:          &digitalTwinsInstanceClient,
		Endpoints:                     &endpointsClient,
		PrivateEndpoints:              &privateEndpointsClient,
		TimeSeriesDatabaseConnections: &timeSeriesDatabaseConnectionsClient,
	}
}
