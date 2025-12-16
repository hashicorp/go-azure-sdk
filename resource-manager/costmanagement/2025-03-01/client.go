package v2025_03_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2025-03-01/alerts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2025-03-01/budgets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2025-03-01/costallocationruledefinitions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2025-03-01/exports"
	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2025-03-01/openapis"
	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2025-03-01/scheduledactionoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2025-03-01/scheduledactions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2025-03-01/settings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2025-03-01/viewoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2025-03-01/views"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Alerts                        *alerts.AlertsClient
	Budgets                       *budgets.BudgetsClient
	CostAllocationRuleDefinitions *costallocationruledefinitions.CostAllocationRuleDefinitionsClient
	Exports                       *exports.ExportsClient
	Openapis                      *openapis.OpenapisClient
	ScheduledActionOperationGroup *scheduledactionoperationgroup.ScheduledActionOperationGroupClient
	ScheduledActions              *scheduledactions.ScheduledActionsClient
	Settings                      *settings.SettingsClient
	ViewOperationGroup            *viewoperationgroup.ViewOperationGroupClient
	Views                         *views.ViewsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	alertsClient, err := alerts.NewAlertsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Alerts client: %+v", err)
	}
	configureFunc(alertsClient.Client)

	budgetsClient, err := budgets.NewBudgetsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Budgets client: %+v", err)
	}
	configureFunc(budgetsClient.Client)

	costAllocationRuleDefinitionsClient, err := costallocationruledefinitions.NewCostAllocationRuleDefinitionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CostAllocationRuleDefinitions client: %+v", err)
	}
	configureFunc(costAllocationRuleDefinitionsClient.Client)

	exportsClient, err := exports.NewExportsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Exports client: %+v", err)
	}
	configureFunc(exportsClient.Client)

	openapisClient, err := openapis.NewOpenapisClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Openapis client: %+v", err)
	}
	configureFunc(openapisClient.Client)

	scheduledActionOperationGroupClient, err := scheduledactionoperationgroup.NewScheduledActionOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ScheduledActionOperationGroup client: %+v", err)
	}
	configureFunc(scheduledActionOperationGroupClient.Client)

	scheduledActionsClient, err := scheduledactions.NewScheduledActionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ScheduledActions client: %+v", err)
	}
	configureFunc(scheduledActionsClient.Client)

	settingsClient, err := settings.NewSettingsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Settings client: %+v", err)
	}
	configureFunc(settingsClient.Client)

	viewOperationGroupClient, err := viewoperationgroup.NewViewOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ViewOperationGroup client: %+v", err)
	}
	configureFunc(viewOperationGroupClient.Client)

	viewsClient, err := views.NewViewsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Views client: %+v", err)
	}
	configureFunc(viewsClient.Client)

	return &Client{
		Alerts:                        alertsClient,
		Budgets:                       budgetsClient,
		CostAllocationRuleDefinitions: costAllocationRuleDefinitionsClient,
		Exports:                       exportsClient,
		Openapis:                      openapisClient,
		ScheduledActionOperationGroup: scheduledActionOperationGroupClient,
		ScheduledActions:              scheduledActionsClient,
		Settings:                      settingsClient,
		ViewOperationGroup:            viewOperationGroupClient,
		Views:                         viewsClient,
	}, nil
}
