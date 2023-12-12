package v2023_10_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2023-10-01/metricdefinitions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2023-10-01/metrics"
)

type Client struct {
	MetricDefinitions *metricdefinitions.MetricDefinitionsClient
	Metrics           *metrics.MetricsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	metricDefinitionsClient := metricdefinitions.NewMetricDefinitionsClientWithBaseURI(endpoint)
	configureAuthFunc(&metricDefinitionsClient.Client)

	metricsClient := metrics.NewMetricsClientWithBaseURI(endpoint)
	configureAuthFunc(&metricsClient.Client)

	return Client{
		MetricDefinitions: &metricDefinitionsClient,
		Metrics:           &metricsClient,
	}
}
