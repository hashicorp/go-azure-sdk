// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2022_10_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2022-10-01/autoscaleapis"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2022-10-01/autoscalesettings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2022-10-01/metrics"
)

type Client struct {
	AutoScaleSettings *autoscalesettings.AutoScaleSettingsClient
	AutoscaleAPIs     *autoscaleapis.AutoscaleAPIsClient
	Metrics           *metrics.MetricsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	autoScaleSettingsClient := autoscalesettings.NewAutoScaleSettingsClientWithBaseURI(endpoint)
	configureAuthFunc(&autoScaleSettingsClient.Client)

	autoscaleAPIsClient := autoscaleapis.NewAutoscaleAPIsClientWithBaseURI(endpoint)
	configureAuthFunc(&autoscaleAPIsClient.Client)

	metricsClient := metrics.NewMetricsClientWithBaseURI(endpoint)
	configureAuthFunc(&metricsClient.Client)

	return Client{
		AutoScaleSettings: &autoScaleSettingsClient,
		AutoscaleAPIs:     &autoscaleAPIsClient,
		Metrics:           &metricsClient,
	}
}
