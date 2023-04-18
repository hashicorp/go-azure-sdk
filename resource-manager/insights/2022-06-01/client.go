package v2022_06_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2022-06-01/actiongroupsapis"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2022-06-01/datacollectionendpoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2022-06-01/datacollectionruleassociations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2022-06-01/datacollectionrules"
)

type Client struct {
	ActionGroupsAPIs               *actiongroupsapis.ActionGroupsAPIsClient
	DataCollectionEndpoints        *datacollectionendpoints.DataCollectionEndpointsClient
	DataCollectionRuleAssociations *datacollectionruleassociations.DataCollectionRuleAssociationsClient
	DataCollectionRules            *datacollectionrules.DataCollectionRulesClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	actionGroupsAPIsClient := actiongroupsapis.NewActionGroupsAPIsClientWithBaseURI(endpoint)
	configureAuthFunc(&actionGroupsAPIsClient.Client)

	dataCollectionEndpointsClient := datacollectionendpoints.NewDataCollectionEndpointsClientWithBaseURI(endpoint)
	configureAuthFunc(&dataCollectionEndpointsClient.Client)

	dataCollectionRuleAssociationsClient := datacollectionruleassociations.NewDataCollectionRuleAssociationsClientWithBaseURI(endpoint)
	configureAuthFunc(&dataCollectionRuleAssociationsClient.Client)

	dataCollectionRulesClient := datacollectionrules.NewDataCollectionRulesClientWithBaseURI(endpoint)
	configureAuthFunc(&dataCollectionRulesClient.Client)

	return Client{
		ActionGroupsAPIs:               &actionGroupsAPIsClient,
		DataCollectionEndpoints:        &dataCollectionEndpointsClient,
		DataCollectionRuleAssociations: &dataCollectionRuleAssociationsClient,
		DataCollectionRules:            &dataCollectionRulesClient,
	}
}
