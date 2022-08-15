package v2021_05_01_preview

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2021-05-01-preview/autoscaleapis"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2021-05-01-preview/autoscalesettings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2021-05-01-preview/diagnosticsettings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2021-05-01-preview/diagnosticsettingscategories"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2021-05-01-preview/managementgroupdiagnosticsettings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2021-05-01-preview/metrics"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2021-05-01-preview/subscriptiondiagnosticsettings"
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

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	autoScaleSettingsClient := autoscalesettings.NewAutoScaleSettingsClientWithBaseURI(endpoint)
	configureAuthFunc(&autoScaleSettingsClient.Client)

	autoscaleAPIsClient := autoscaleapis.NewAutoscaleAPIsClientWithBaseURI(endpoint)
	configureAuthFunc(&autoscaleAPIsClient.Client)

	diagnosticSettingsCategoriesClient := diagnosticsettingscategories.NewDiagnosticSettingsCategoriesClientWithBaseURI(endpoint)
	configureAuthFunc(&diagnosticSettingsCategoriesClient.Client)

	diagnosticSettingsClient := diagnosticsettings.NewDiagnosticSettingsClientWithBaseURI(endpoint)
	configureAuthFunc(&diagnosticSettingsClient.Client)

	managementGroupDiagnosticSettingsClient := managementgroupdiagnosticsettings.NewManagementGroupDiagnosticSettingsClientWithBaseURI(endpoint)
	configureAuthFunc(&managementGroupDiagnosticSettingsClient.Client)

	metricsClient := metrics.NewMetricsClientWithBaseURI(endpoint)
	configureAuthFunc(&metricsClient.Client)

	subscriptionDiagnosticSettingsClient := subscriptiondiagnosticsettings.NewSubscriptionDiagnosticSettingsClientWithBaseURI(endpoint)
	configureAuthFunc(&subscriptionDiagnosticSettingsClient.Client)

	return Client{
		AutoScaleSettings:                 &autoScaleSettingsClient,
		AutoscaleAPIs:                     &autoscaleAPIsClient,
		DiagnosticSettings:                &diagnosticSettingsClient,
		DiagnosticSettingsCategories:      &diagnosticSettingsCategoriesClient,
		ManagementGroupDiagnosticSettings: &managementGroupDiagnosticSettingsClient,
		Metrics:                           &metricsClient,
		SubscriptionDiagnosticSettings:    &subscriptionDiagnosticSettingsClient,
	}
}
