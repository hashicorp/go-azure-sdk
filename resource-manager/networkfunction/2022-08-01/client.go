// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2022_08_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/networkfunction/2022-08-01/azuretrafficcollectors"
	"github.com/hashicorp/go-azure-sdk/resource-manager/networkfunction/2022-08-01/collectorpolicies"
)

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
