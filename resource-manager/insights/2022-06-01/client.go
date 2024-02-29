package v2022_06_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2022-06-01/actiongroupsapis"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2022-06-01/datacollectionendpoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2022-06-01/datacollectionruleassociations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2022-06-01/datacollectionrules"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	ActionGroupsAPIs               *actiongroupsapis.ActionGroupsAPIsClient
	DataCollectionEndpoints        *datacollectionendpoints.DataCollectionEndpointsClient
	DataCollectionRuleAssociations *datacollectionruleassociations.DataCollectionRuleAssociationsClient
	DataCollectionRules            *datacollectionrules.DataCollectionRulesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	actionGroupsAPIsClient, err := actiongroupsapis.NewActionGroupsAPIsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ActionGroupsAPIs client: %+v", err)
	}
	configureFunc(actionGroupsAPIsClient.Client)

	dataCollectionEndpointsClient, err := datacollectionendpoints.NewDataCollectionEndpointsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DataCollectionEndpoints client: %+v", err)
	}
	configureFunc(dataCollectionEndpointsClient.Client)

	dataCollectionRuleAssociationsClient, err := datacollectionruleassociations.NewDataCollectionRuleAssociationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DataCollectionRuleAssociations client: %+v", err)
	}
	configureFunc(dataCollectionRuleAssociationsClient.Client)

	dataCollectionRulesClient, err := datacollectionrules.NewDataCollectionRulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DataCollectionRules client: %+v", err)
	}
	configureFunc(dataCollectionRulesClient.Client)

	return &Client{
		ActionGroupsAPIs:               actionGroupsAPIsClient,
		DataCollectionEndpoints:        dataCollectionEndpointsClient,
		DataCollectionRuleAssociations: dataCollectionRuleAssociationsClient,
		DataCollectionRules:            dataCollectionRulesClient,
	}, nil
}
