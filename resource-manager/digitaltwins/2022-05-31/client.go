package v2022_05_31

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/digitaltwins/2022-05-31/checknameavailability"
	"github.com/hashicorp/go-azure-sdk/resource-manager/digitaltwins/2022-05-31/digitaltwinsinstance"
	"github.com/hashicorp/go-azure-sdk/resource-manager/digitaltwins/2022-05-31/endpoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/digitaltwins/2022-05-31/privateendpoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/digitaltwins/2022-05-31/timeseriesdatabaseconnections"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

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
