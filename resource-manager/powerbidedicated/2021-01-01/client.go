package v2021_01_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/powerbidedicated/2021-01-01/autoscalevcores"
	"github.com/hashicorp/go-azure-sdk/resource-manager/powerbidedicated/2021-01-01/capacities"
	"github.com/hashicorp/go-azure-sdk/resource-manager/powerbidedicated/2021-01-01/powerbidedicated"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AutoScaleVCores  *autoscalevcores.AutoScaleVCoresClient
	Capacities       *capacities.CapacitiesClient
	PowerBIDedicated *powerbidedicated.PowerBIDedicatedClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	autoScaleVCoresClient, err := autoscalevcores.NewAutoScaleVCoresClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AutoScaleVCores client: %+v", err)
	}
	configureFunc(autoScaleVCoresClient.Client)

	capacitiesClient, err := capacities.NewCapacitiesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Capacities client: %+v", err)
	}
	configureFunc(capacitiesClient.Client)

	powerBIDedicatedClient, err := powerbidedicated.NewPowerBIDedicatedClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PowerBIDedicated client: %+v", err)
	}
	configureFunc(powerBIDedicatedClient.Client)

	return &Client{
		AutoScaleVCores:  autoScaleVCoresClient,
		Capacities:       capacitiesClient,
		PowerBIDedicated: powerBIDedicatedClient,
	}, nil
}
