package v2021_04_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2021-04-01/datacollectionendpoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2021-04-01/datacollectionruleassociations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2021-04-01/datacollectionrules"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Client struct {
	DataCollectionEndpoints        *datacollectionendpoints.DataCollectionEndpointsClient
	DataCollectionRuleAssociations *datacollectionruleassociations.DataCollectionRuleAssociationsClient
	DataCollectionRules            *datacollectionrules.DataCollectionRulesClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	dataCollectionEndpointsClient := datacollectionendpoints.NewDataCollectionEndpointsClientWithBaseURI(endpoint)
	configureAuthFunc(&dataCollectionEndpointsClient.Client)

	dataCollectionRuleAssociationsClient := datacollectionruleassociations.NewDataCollectionRuleAssociationsClientWithBaseURI(endpoint)
	configureAuthFunc(&dataCollectionRuleAssociationsClient.Client)

	dataCollectionRulesClient := datacollectionrules.NewDataCollectionRulesClientWithBaseURI(endpoint)
	configureAuthFunc(&dataCollectionRulesClient.Client)

	return Client{
		DataCollectionEndpoints:        &dataCollectionEndpointsClient,
		DataCollectionRuleAssociations: &dataCollectionRuleAssociationsClient,
		DataCollectionRules:            &dataCollectionRulesClient,
	}
}
