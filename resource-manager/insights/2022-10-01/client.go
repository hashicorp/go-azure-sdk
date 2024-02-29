package v2022_10_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2022-10-01/autoscaleapis"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2022-10-01/autoscalesettings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2022-10-01/metrics"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AutoScaleSettings *autoscalesettings.AutoScaleSettingsClient
	AutoscaleAPIs     *autoscaleapis.AutoscaleAPIsClient
	Metrics           *metrics.MetricsClient
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

	metricsClient, err := metrics.NewMetricsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Metrics client: %+v", err)
	}
	configureFunc(metricsClient.Client)

	return &Client{
		AutoScaleSettings: autoScaleSettingsClient,
		AutoscaleAPIs:     autoscaleAPIsClient,
		Metrics:           metricsClient,
	}, nil
}
