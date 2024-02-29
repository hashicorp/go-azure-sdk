package v2015_04_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2015-04-01/activitylogs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2015-04-01/autoscaleapis"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2015-04-01/autoscalesettings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2015-04-01/eventcategories"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2015-04-01/tenantactivitylogs"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	ActivityLogs       *activitylogs.ActivityLogsClient
	AutoScaleSettings  *autoscalesettings.AutoScaleSettingsClient
	AutoscaleAPIs      *autoscaleapis.AutoscaleAPIsClient
	EventCategories    *eventcategories.EventCategoriesClient
	TenantActivityLogs *tenantactivitylogs.TenantActivityLogsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	activityLogsClient, err := activitylogs.NewActivityLogsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ActivityLogs client: %+v", err)
	}
	configureFunc(activityLogsClient.Client)

	autoScaleSettingsClient, err := autoscalesettings.NewAutoScaleSettingsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AutoScaleSettings client: %+v", err)
	}
	configureFunc(autoScaleSettingsClient.Client)

	autoscaleAPIsClient, err := autoscaleapis.NewAutoscaleAPIsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AutoscaleAPIs client: %+v", err)
	}
	configureFunc(autoscaleAPIsClient.Client)

	eventCategoriesClient, err := eventcategories.NewEventCategoriesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EventCategories client: %+v", err)
	}
	configureFunc(eventCategoriesClient.Client)

	tenantActivityLogsClient, err := tenantactivitylogs.NewTenantActivityLogsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TenantActivityLogs client: %+v", err)
	}
	configureFunc(tenantActivityLogsClient.Client)

	return &Client{
		ActivityLogs:       activityLogsClient,
		AutoScaleSettings:  autoScaleSettingsClient,
		AutoscaleAPIs:      autoscaleAPIsClient,
		EventCategories:    eventCategoriesClient,
		TenantActivityLogs: tenantActivityLogsClient,
	}, nil
}
