// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2022_10_01

import (
	"github.com/Azure/go-autorest/autorest"
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
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

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

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	alertsClient := alerts.NewAlertsClientWithBaseURI(endpoint)
	configureAuthFunc(&alertsClient.Client)

	benefitRecommendationsClient := benefitrecommendations.NewBenefitRecommendationsClientWithBaseURI(endpoint)
	configureAuthFunc(&benefitRecommendationsClient.Client)

	benefitUtilizationSummariesClient := benefitutilizationsummaries.NewBenefitUtilizationSummariesClientWithBaseURI(endpoint)
	configureAuthFunc(&benefitUtilizationSummariesClient.Client)

	costDetailsClient := costdetails.NewCostDetailsClientWithBaseURI(endpoint)
	configureAuthFunc(&costDetailsClient.Client)

	dimensionsClient := dimensions.NewDimensionsClientWithBaseURI(endpoint)
	configureAuthFunc(&dimensionsClient.Client)

	exportsClient := exports.NewExportsClientWithBaseURI(endpoint)
	configureAuthFunc(&exportsClient.Client)

	forecastClient := forecast.NewForecastClientWithBaseURI(endpoint)
	configureAuthFunc(&forecastClient.Client)

	priceSheetsClient := pricesheets.NewPriceSheetsClientWithBaseURI(endpoint)
	configureAuthFunc(&priceSheetsClient.Client)

	queryClient := query.NewQueryClientWithBaseURI(endpoint)
	configureAuthFunc(&queryClient.Client)

	reservedInstancesClient := reservedinstances.NewReservedInstancesClientWithBaseURI(endpoint)
	configureAuthFunc(&reservedInstancesClient.Client)

	scheduledActionsClient := scheduledactions.NewScheduledActionsClientWithBaseURI(endpoint)
	configureAuthFunc(&scheduledActionsClient.Client)

	usageDetailsClient := usagedetails.NewUsageDetailsClientWithBaseURI(endpoint)
	configureAuthFunc(&usageDetailsClient.Client)

	viewsClient := views.NewViewsClientWithBaseURI(endpoint)
	configureAuthFunc(&viewsClient.Client)

	return Client{
		Alerts:                      &alertsClient,
		BenefitRecommendations:      &benefitRecommendationsClient,
		BenefitUtilizationSummaries: &benefitUtilizationSummariesClient,
		CostDetails:                 &costDetailsClient,
		Dimensions:                  &dimensionsClient,
		Exports:                     &exportsClient,
		Forecast:                    &forecastClient,
		PriceSheets:                 &priceSheetsClient,
		Query:                       &queryClient,
		ReservedInstances:           &reservedInstancesClient,
		ScheduledActions:            &scheduledActionsClient,
		UsageDetails:                &usageDetailsClient,
		Views:                       &viewsClient,
	}
}
