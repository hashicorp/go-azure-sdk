package v2021_05_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2021-05-01-preview/autoscaleapis"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2021-05-01-preview/autoscalesettings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2021-05-01-preview/diagnosticsettings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2021-05-01-preview/diagnosticsettingscategories"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2021-05-01-preview/managementgroupdiagnosticsettings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2021-05-01-preview/metrics"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2021-05-01-preview/subscriptiondiagnosticsettings"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AutoScaleSettings                 *autoscalesettings.AutoScaleSettingsClient
	AutoscaleAPIs                     *autoscaleapis.AutoscaleAPIsClient
	DiagnosticSettings                *diagnosticsettings.DiagnosticSettingsClient
	DiagnosticSettingsCategories      *diagnosticsettingscategories.DiagnosticSettingsCategoriesClient
	ManagementGroupDiagnosticSettings *managementgroupdiagnosticsettings.ManagementGroupDiagnosticSettingsClient
	Metrics                           *metrics.MetricsClient
	SubscriptionDiagnosticSettings    *subscriptiondiagnosticsettings.SubscriptionDiagnosticSettingsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
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

	diagnosticSettingsCategoriesClient, err := diagnosticsettingscategories.NewDiagnosticSettingsCategoriesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DiagnosticSettingsCategories client: %+v", err)
	}
	configureFunc(diagnosticSettingsCategoriesClient.Client)

	diagnosticSettingsClient, err := diagnosticsettings.NewDiagnosticSettingsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DiagnosticSettings client: %+v", err)
	}
	configureFunc(diagnosticSettingsClient.Client)

	managementGroupDiagnosticSettingsClient, err := managementgroupdiagnosticsettings.NewManagementGroupDiagnosticSettingsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagementGroupDiagnosticSettings client: %+v", err)
	}
	configureFunc(managementGroupDiagnosticSettingsClient.Client)

	metricsClient, err := metrics.NewMetricsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Metrics client: %+v", err)
	}
	configureFunc(metricsClient.Client)

	subscriptionDiagnosticSettingsClient, err := subscriptiondiagnosticsettings.NewSubscriptionDiagnosticSettingsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SubscriptionDiagnosticSettings client: %+v", err)
	}
	configureFunc(subscriptionDiagnosticSettingsClient.Client)

	return &Client{
		AutoScaleSettings:                 autoScaleSettingsClient,
		AutoscaleAPIs:                     autoscaleAPIsClient,
		DiagnosticSettings:                diagnosticSettingsClient,
		DiagnosticSettingsCategories:      diagnosticSettingsCategoriesClient,
		ManagementGroupDiagnosticSettings: managementGroupDiagnosticSettingsClient,
		Metrics:                           metricsClient,
		SubscriptionDiagnosticSettings:    subscriptionDiagnosticSettingsClient,
	}, nil
}
