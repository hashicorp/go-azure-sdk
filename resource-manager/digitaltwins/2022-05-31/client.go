package v2022_05_31

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/digitaltwins/2022-05-31/checknameavailability"
	"github.com/hashicorp/go-azure-sdk/resource-manager/digitaltwins/2022-05-31/digitaltwinsinstance"
	"github.com/hashicorp/go-azure-sdk/resource-manager/digitaltwins/2022-05-31/endpoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/digitaltwins/2022-05-31/privateendpoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/digitaltwins/2022-05-31/timeseriesdatabaseconnections"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	CheckNameAvailability         *checknameavailability.CheckNameAvailabilityClient
	DigitalTwinsInstance          *digitaltwinsinstance.DigitalTwinsInstanceClient
	Endpoints                     *endpoints.EndpointsClient
	PrivateEndpoints              *privateendpoints.PrivateEndpointsClient
	TimeSeriesDatabaseConnections *timeseriesdatabaseconnections.TimeSeriesDatabaseConnectionsClient
}

func NewClientWithBaseURI(api environments.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	checkNameAvailabilityClient, err := checknameavailability.NewCheckNameAvailabilityClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building CheckNameAvailability client: %+v", err)
	}
	configureFunc(checkNameAvailabilityClient.Client)

	digitalTwinsInstanceClient, err := digitaltwinsinstance.NewDigitalTwinsInstanceClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building DigitalTwinsInstance client: %+v", err)
	}
	configureFunc(digitalTwinsInstanceClient.Client)

	endpointsClient, err := endpoints.NewEndpointsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Endpoints client: %+v", err)
	}
	configureFunc(endpointsClient.Client)

	privateEndpointsClient, err := privateendpoints.NewPrivateEndpointsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building PrivateEndpoints client: %+v", err)
	}
	configureFunc(privateEndpointsClient.Client)

	timeSeriesDatabaseConnectionsClient, err := timeseriesdatabaseconnections.NewTimeSeriesDatabaseConnectionsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building TimeSeriesDatabaseConnections client: %+v", err)
	}
	configureFunc(timeSeriesDatabaseConnectionsClient.Client)

	return &Client{
		CheckNameAvailability:         checkNameAvailabilityClient,
		DigitalTwinsInstance:          digitalTwinsInstanceClient,
		Endpoints:                     endpointsClient,
		PrivateEndpoints:              privateEndpointsClient,
		TimeSeriesDatabaseConnections: timeSeriesDatabaseConnectionsClient,
	}, nil
}
