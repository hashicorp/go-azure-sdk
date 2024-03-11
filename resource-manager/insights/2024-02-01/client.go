package v2024_02_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2024-02-01/metricdefinitions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2024-02-01/metricnamespaces"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2024-02-01/metrics"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	MetricDefinitions *metricdefinitions.MetricDefinitionsClient
	MetricNamespaces  *metricnamespaces.MetricNamespacesClient
	Metrics           *metrics.MetricsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	metricDefinitionsClient, err := metricdefinitions.NewMetricDefinitionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MetricDefinitions client: %+v", err)
	}
	configureFunc(metricDefinitionsClient.Client)

	metricNamespacesClient, err := metricnamespaces.NewMetricNamespacesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MetricNamespaces client: %+v", err)
	}
	configureFunc(metricNamespacesClient.Client)

	metricsClient, err := metrics.NewMetricsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Metrics client: %+v", err)
	}
	configureFunc(metricsClient.Client)

	return &Client{
		MetricDefinitions: metricDefinitionsClient,
		MetricNamespaces:  metricNamespacesClient,
		Metrics:           metricsClient,
	}, nil
}
