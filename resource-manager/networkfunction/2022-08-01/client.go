package v2022_08_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/networkfunction/2022-08-01/azuretrafficcollectors"
	"github.com/hashicorp/go-azure-sdk/resource-manager/networkfunction/2022-08-01/collectorpolicies"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Client struct {
	AzureTrafficCollectors *azuretrafficcollectors.AzureTrafficCollectorsClient
	CollectorPolicies      *collectorpolicies.CollectorPoliciesClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	azureTrafficCollectorsClient := azuretrafficcollectors.NewAzureTrafficCollectorsClientWithBaseURI(endpoint)
	configureAuthFunc(&azureTrafficCollectorsClient.Client)

	collectorPoliciesClient := collectorpolicies.NewCollectorPoliciesClientWithBaseURI(endpoint)
	configureAuthFunc(&collectorPoliciesClient.Client)

	return Client{
		AzureTrafficCollectors: &azureTrafficCollectorsClient,
		CollectorPolicies:      &collectorPoliciesClient,
	}
}
