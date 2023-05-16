package v2022_08_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/networkfunction/2022-08-01/azuretrafficcollectors"
	"github.com/hashicorp/go-azure-sdk/resource-manager/networkfunction/2022-08-01/collectorpolicies"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AzureTrafficCollectors *azuretrafficcollectors.AzureTrafficCollectorsClient
	CollectorPolicies      *collectorpolicies.CollectorPoliciesClient
}

func NewClientWithBaseURI(api environments.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	azureTrafficCollectorsClient, err := azuretrafficcollectors.NewAzureTrafficCollectorsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building AzureTrafficCollectors client: %+v", err)
	}
	configureFunc(azureTrafficCollectorsClient.Client)

	collectorPoliciesClient, err := collectorpolicies.NewCollectorPoliciesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building CollectorPolicies client: %+v", err)
	}
	configureFunc(collectorPoliciesClient.Client)

	return &Client{
		AzureTrafficCollectors: azureTrafficCollectorsClient,
		CollectorPolicies:      collectorPoliciesClient,
	}, nil
}
