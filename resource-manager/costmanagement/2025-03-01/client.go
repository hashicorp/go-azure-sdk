package v2025_03_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2025-03-01/alerts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2025-03-01/benefitrecommendations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2025-03-01/benefitutilizationsummaries"
	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2025-03-01/benefitutilizationsummariesasync"
	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2025-03-01/budgets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2025-03-01/costallocationrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2025-03-01/costdetails"
	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2025-03-01/dimensions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2025-03-01/exports"
	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2025-03-01/forecast"
	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2025-03-01/pricesheets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2025-03-01/query"
	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2025-03-01/reservedinstances"
	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2025-03-01/scheduledactions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2025-03-01/settings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2025-03-01/usagedetails"
	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2025-03-01/views"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Alerts                           *alerts.AlertsClient
	BenefitRecommendations           *benefitrecommendations.BenefitRecommendationsClient
	BenefitUtilizationSummaries      *benefitutilizationsummaries.BenefitUtilizationSummariesClient
	BenefitUtilizationSummariesAsync *benefitutilizationsummariesasync.BenefitUtilizationSummariesAsyncClient
	Budgets                          *budgets.BudgetsClient
	CostAllocationRules              *costallocationrules.CostAllocationRulesClient
	CostDetails                      *costdetails.CostDetailsClient
	Dimensions                       *dimensions.DimensionsClient
	Exports                          *exports.ExportsClient
	Forecast                         *forecast.ForecastClient
	PriceSheets                      *pricesheets.PriceSheetsClient
	Query                            *query.QueryClient
	ReservedInstances                *reservedinstances.ReservedInstancesClient
	ScheduledActions                 *scheduledactions.ScheduledActionsClient
	Settings                         *settings.SettingsClient
	UsageDetails                     *usagedetails.UsageDetailsClient
	Views                            *views.ViewsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	alertsClient, err := alerts.NewAlertsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Alerts client: %+v", err)
	}
	configureFunc(alertsClient.Client)

	benefitRecommendationsClient, err := benefitrecommendations.NewBenefitRecommendationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BenefitRecommendations client: %+v", err)
	}
	configureFunc(benefitRecommendationsClient.Client)

	benefitUtilizationSummariesAsyncClient, err := benefitutilizationsummariesasync.NewBenefitUtilizationSummariesAsyncClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BenefitUtilizationSummariesAsync client: %+v", err)
	}
	configureFunc(benefitUtilizationSummariesAsyncClient.Client)

	benefitUtilizationSummariesClient, err := benefitutilizationsummaries.NewBenefitUtilizationSummariesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BenefitUtilizationSummaries client: %+v", err)
	}
	configureFunc(benefitUtilizationSummariesClient.Client)

	budgetsClient, err := budgets.NewBudgetsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Budgets client: %+v", err)
	}
	configureFunc(budgetsClient.Client)

	costAllocationRulesClient, err := costallocationrules.NewCostAllocationRulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CostAllocationRules client: %+v", err)
	}
	configureFunc(costAllocationRulesClient.Client)

	costDetailsClient, err := costdetails.NewCostDetailsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CostDetails client: %+v", err)
	}
	configureFunc(costDetailsClient.Client)

	dimensionsClient, err := dimensions.NewDimensionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Dimensions client: %+v", err)
	}
	configureFunc(dimensionsClient.Client)

	exportsClient, err := exports.NewExportsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Exports client: %+v", err)
	}
	configureFunc(exportsClient.Client)

	forecastClient, err := forecast.NewForecastClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Forecast client: %+v", err)
	}
	configureFunc(forecastClient.Client)

	priceSheetsClient, err := pricesheets.NewPriceSheetsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PriceSheets client: %+v", err)
	}
	configureFunc(priceSheetsClient.Client)

	queryClient, err := query.NewQueryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Query client: %+v", err)
	}
	configureFunc(queryClient.Client)

	reservedInstancesClient, err := reservedinstances.NewReservedInstancesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ReservedInstances client: %+v", err)
	}
	configureFunc(reservedInstancesClient.Client)

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

	usageDetailsClient, err := usagedetails.NewUsageDetailsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UsageDetails client: %+v", err)
	}
	configureFunc(usageDetailsClient.Client)

	viewsClient, err := views.NewViewsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Views client: %+v", err)
	}
	configureFunc(viewsClient.Client)

	return &Client{
		Alerts:                           alertsClient,
		BenefitRecommendations:           benefitRecommendationsClient,
		BenefitUtilizationSummaries:      benefitUtilizationSummariesClient,
		BenefitUtilizationSummariesAsync: benefitUtilizationSummariesAsyncClient,
		Budgets:                          budgetsClient,
		CostAllocationRules:              costAllocationRulesClient,
		CostDetails:                      costDetailsClient,
		Dimensions:                       dimensionsClient,
		Exports:                          exportsClient,
		Forecast:                         forecastClient,
		PriceSheets:                      priceSheetsClient,
		Query:                            queryClient,
		ReservedInstances:                reservedInstancesClient,
		ScheduledActions:                 scheduledActionsClient,
		Settings:                         settingsClient,
		UsageDetails:                     usageDetailsClient,
		Views:                            viewsClient,
	}, nil
}
