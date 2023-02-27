// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2021_10_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2021-10-01/alerts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2021-10-01/dimensions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2021-10-01/exports"
	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2021-10-01/forecast"
	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2021-10-01/query"
	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2021-10-01/reservedinstances"
	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2021-10-01/usagedetails"
	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2021-10-01/views"
)

type Client struct {
	Alerts            *alerts.AlertsClient
	Dimensions        *dimensions.DimensionsClient
	Exports           *exports.ExportsClient
	Forecast          *forecast.ForecastClient
	Query             *query.QueryClient
	ReservedInstances *reservedinstances.ReservedInstancesClient
	UsageDetails      *usagedetails.UsageDetailsClient
	Views             *views.ViewsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	alertsClient := alerts.NewAlertsClientWithBaseURI(endpoint)
	configureAuthFunc(&alertsClient.Client)

	dimensionsClient := dimensions.NewDimensionsClientWithBaseURI(endpoint)
	configureAuthFunc(&dimensionsClient.Client)

	exportsClient := exports.NewExportsClientWithBaseURI(endpoint)
	configureAuthFunc(&exportsClient.Client)

	forecastClient := forecast.NewForecastClientWithBaseURI(endpoint)
	configureAuthFunc(&forecastClient.Client)

	queryClient := query.NewQueryClientWithBaseURI(endpoint)
	configureAuthFunc(&queryClient.Client)

	reservedInstancesClient := reservedinstances.NewReservedInstancesClientWithBaseURI(endpoint)
	configureAuthFunc(&reservedInstancesClient.Client)

	usageDetailsClient := usagedetails.NewUsageDetailsClientWithBaseURI(endpoint)
	configureAuthFunc(&usageDetailsClient.Client)

	viewsClient := views.NewViewsClientWithBaseURI(endpoint)
	configureAuthFunc(&viewsClient.Client)

	return Client{
		Alerts:            &alertsClient,
		Dimensions:        &dimensionsClient,
		Exports:           &exportsClient,
		Forecast:          &forecastClient,
		Query:             &queryClient,
		ReservedInstances: &reservedInstancesClient,
		UsageDetails:      &usageDetailsClient,
		Views:             &viewsClient,
	}
}
