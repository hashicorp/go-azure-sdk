package v2022_10_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2022-10-01/alerts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2022-10-01/benefitrecommendations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2022-10-01/benefitutilizationsummaries"
	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2022-10-01/costdetails"
	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2022-10-01/dimensions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2022-10-01/exports"
	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2022-10-01/forecast"
	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2022-10-01/pricesheets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2022-10-01/query"
	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2022-10-01/reservedinstances"
	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2022-10-01/scheduledactions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2022-10-01/usagedetails"
	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2022-10-01/views"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Alerts                      *alerts.AlertsClient
	BenefitRecommendations      *benefitrecommendations.BenefitRecommendationsClient
	BenefitUtilizationSummaries *benefitutilizationsummaries.BenefitUtilizationSummariesClient
	CostDetails                 *costdetails.CostDetailsClient
	Dimensions                  *dimensions.DimensionsClient
	Exports                     *exports.ExportsClient
	Forecast                    *forecast.ForecastClient
	PriceSheets                 *pricesheets.PriceSheetsClient
	Query                       *query.QueryClient
	ReservedInstances           *reservedinstances.ReservedInstancesClient
	ScheduledActions            *scheduledactions.ScheduledActionsClient
	UsageDetails                *usagedetails.UsageDetailsClient
	Views                       *views.ViewsClient
}

func NewClientWithBaseURI(api environments.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	alertsClient, err := alerts.NewAlertsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Alerts client: %+v", err)
	}
	configureFunc(alertsClient.Client)

	benefitRecommendationsClient, err := benefitrecommendations.NewBenefitRecommendationsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building BenefitRecommendations client: %+v", err)
	}
	configureFunc(benefitRecommendationsClient.Client)

	benefitUtilizationSummariesClient, err := benefitutilizationsummaries.NewBenefitUtilizationSummariesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building BenefitUtilizationSummaries client: %+v", err)
	}
	configureFunc(benefitUtilizationSummariesClient.Client)

	costDetailsClient, err := costdetails.NewCostDetailsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building CostDetails client: %+v", err)
	}
	configureFunc(costDetailsClient.Client)

	dimensionsClient, err := dimensions.NewDimensionsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Dimensions client: %+v", err)
	}
	configureFunc(dimensionsClient.Client)

	exportsClient, err := exports.NewExportsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Exports client: %+v", err)
	}
	configureFunc(exportsClient.Client)

	forecastClient, err := forecast.NewForecastClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Forecast client: %+v", err)
	}
	configureFunc(forecastClient.Client)

	priceSheetsClient, err := pricesheets.NewPriceSheetsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building PriceSheets client: %+v", err)
	}
	configureFunc(priceSheetsClient.Client)

	queryClient, err := query.NewQueryClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Query client: %+v", err)
	}
	configureFunc(queryClient.Client)

	reservedInstancesClient, err := reservedinstances.NewReservedInstancesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ReservedInstances client: %+v", err)
	}
	configureFunc(reservedInstancesClient.Client)

	scheduledActionsClient, err := scheduledactions.NewScheduledActionsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ScheduledActions client: %+v", err)
	}
	configureFunc(scheduledActionsClient.Client)

	usageDetailsClient, err := usagedetails.NewUsageDetailsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building UsageDetails client: %+v", err)
	}
	configureFunc(usageDetailsClient.Client)

	viewsClient, err := views.NewViewsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Views client: %+v", err)
	}
	configureFunc(viewsClient.Client)

	return &Client{
		Alerts:                      alertsClient,
		BenefitRecommendations:      benefitRecommendationsClient,
		BenefitUtilizationSummaries: benefitUtilizationSummariesClient,
		CostDetails:                 costDetailsClient,
		Dimensions:                  dimensionsClient,
		Exports:                     exportsClient,
		Forecast:                    forecastClient,
		PriceSheets:                 priceSheetsClient,
		Query:                       queryClient,
		ReservedInstances:           reservedInstancesClient,
		ScheduledActions:            scheduledActionsClient,
		UsageDetails:                usageDetailsClient,
		Views:                       viewsClient,
	}, nil
}
