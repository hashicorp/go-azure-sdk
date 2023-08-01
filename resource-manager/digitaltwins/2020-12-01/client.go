package v2020_12_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/digitaltwins/2020-12-01/checknameavailability"
	"github.com/hashicorp/go-azure-sdk/resource-manager/digitaltwins/2020-12-01/digitaltwinsinstance"
	"github.com/hashicorp/go-azure-sdk/resource-manager/digitaltwins/2020-12-01/endpoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/digitaltwins/2020-12-01/privateendpoints"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	CheckNameAvailability *checknameavailability.CheckNameAvailabilityClient
	DigitalTwinsInstance  *digitaltwinsinstance.DigitalTwinsInstanceClient
	Endpoints             *endpoints.EndpointsClient
	PrivateEndpoints      *privateendpoints.PrivateEndpointsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	checkNameAvailabilityClient, err := checknameavailability.NewCheckNameAvailabilityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CheckNameAvailability client: %+v", err)
	}
	configureFunc(checkNameAvailabilityClient.Client)

	digitalTwinsInstanceClient, err := digitaltwinsinstance.NewDigitalTwinsInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DigitalTwinsInstance client: %+v", err)
	}
	configureFunc(digitalTwinsInstanceClient.Client)

	endpointsClient, err := endpoints.NewEndpointsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Endpoints client: %+v", err)
	}
	configureFunc(endpointsClient.Client)

	privateEndpointsClient, err := privateendpoints.NewPrivateEndpointsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateEndpoints client: %+v", err)
	}
	configureFunc(privateEndpointsClient.Client)

	return &Client{
		CheckNameAvailability: checkNameAvailabilityClient,
		DigitalTwinsInstance:  digitalTwinsInstanceClient,
		Endpoints:             endpointsClient,
		PrivateEndpoints:      privateEndpointsClient,
	}, nil
}
