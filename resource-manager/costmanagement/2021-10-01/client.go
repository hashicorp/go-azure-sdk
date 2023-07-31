package v2021_10_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2021-10-01/alerts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2021-10-01/dimensions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2021-10-01/exports"
	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2021-10-01/forecast"
	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2021-10-01/query"
	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2021-10-01/reservedinstances"
	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2021-10-01/usagedetails"
	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2021-10-01/views"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
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

func NewClientWithBaseURI(api sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	alertsClient, err := alerts.NewAlertsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Alerts client: %+v", err)
	}
	configureFunc(alertsClient.Client)

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
		Alerts:            alertsClient,
		Dimensions:        dimensionsClient,
		Exports:           exportsClient,
		Forecast:          forecastClient,
		Query:             queryClient,
		ReservedInstances: reservedInstancesClient,
		UsageDetails:      usageDetailsClient,
		Views:             viewsClient,
	}, nil
}
