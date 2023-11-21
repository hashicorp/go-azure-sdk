package v2019_05_05_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/alertsmanagement/2019-05-05-preview/actionrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/alertsmanagement/2019-05-05-preview/alertsmanagements"
	"github.com/hashicorp/go-azure-sdk/resource-manager/alertsmanagement/2019-05-05-preview/smartgroups"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	ActionRules       *actionrules.ActionRulesClient
	AlertsManagements *alertsmanagements.AlertsManagementsClient
	SmartGroups       *smartgroups.SmartGroupsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	actionRulesClient, err := actionrules.NewActionRulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ActionRules client: %+v", err)
	}
	configureFunc(actionRulesClient.Client)

	alertsManagementsClient, err := alertsmanagements.NewAlertsManagementsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AlertsManagements client: %+v", err)
	}
	configureFunc(alertsManagementsClient.Client)

	smartGroupsClient, err := smartgroups.NewSmartGroupsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SmartGroups client: %+v", err)
	}
	configureFunc(smartGroupsClient.Client)

	return &Client{
		ActionRules:       actionRulesClient,
		AlertsManagements: alertsManagementsClient,
		SmartGroups:       smartGroupsClient,
	}, nil
}
