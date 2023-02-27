// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2020_12_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/digitaltwins/2020-12-01/checknameavailability"
	"github.com/hashicorp/go-azure-sdk/resource-manager/digitaltwins/2020-12-01/digitaltwinsinstance"
	"github.com/hashicorp/go-azure-sdk/resource-manager/digitaltwins/2020-12-01/endpoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/digitaltwins/2020-12-01/privateendpoints"
)

type Client struct {
	CheckNameAvailability *checknameavailability.CheckNameAvailabilityClient
	DigitalTwinsInstance  *digitaltwinsinstance.DigitalTwinsInstanceClient
	Endpoints             *endpoints.EndpointsClient
	PrivateEndpoints      *privateendpoints.PrivateEndpointsClient
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

	return Client{
		CheckNameAvailability: &checkNameAvailabilityClient,
		DigitalTwinsInstance:  &digitalTwinsInstanceClient,
		Endpoints:             &endpointsClient,
		PrivateEndpoints:      &privateEndpointsClient,
	}
}
