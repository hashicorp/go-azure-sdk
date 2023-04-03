package v2015_04_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2015-04-01/activitylogs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2015-04-01/autoscaleapis"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2015-04-01/autoscalesettings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2015-04-01/eventcategories"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2015-04-01/tenantactivitylogs"
)

type Client struct {
	ActivityLogs       *activitylogs.ActivityLogsClient
	AutoScaleSettings  *autoscalesettings.AutoScaleSettingsClient
	AutoscaleAPIs      *autoscaleapis.AutoscaleAPIsClient
	EventCategories    *eventcategories.EventCategoriesClient
	TenantActivityLogs *tenantactivitylogs.TenantActivityLogsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	activityLogsClient := activitylogs.NewActivityLogsClientWithBaseURI(endpoint)
	configureAuthFunc(&activityLogsClient.Client)

	autoScaleSettingsClient := autoscalesettings.NewAutoScaleSettingsClientWithBaseURI(endpoint)
	configureAuthFunc(&autoScaleSettingsClient.Client)

	autoscaleAPIsClient := autoscaleapis.NewAutoscaleAPIsClientWithBaseURI(endpoint)
	configureAuthFunc(&autoscaleAPIsClient.Client)

	eventCategoriesClient := eventcategories.NewEventCategoriesClientWithBaseURI(endpoint)
	configureAuthFunc(&eventCategoriesClient.Client)

	tenantActivityLogsClient := tenantactivitylogs.NewTenantActivityLogsClientWithBaseURI(endpoint)
	configureAuthFunc(&tenantActivityLogsClient.Client)

	return Client{
		ActivityLogs:       &activityLogsClient,
		AutoScaleSettings:  &autoScaleSettingsClient,
		AutoscaleAPIs:      &autoscaleAPIsClient,
		EventCategories:    &eventCategoriesClient,
		TenantActivityLogs: &tenantActivityLogsClient,
	}
}
