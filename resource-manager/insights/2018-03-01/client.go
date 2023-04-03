package v2018_03_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2018-03-01/actiongroupsapis"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2018-03-01/metricalerts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2018-03-01/metricalertsstatus"
)

type Client struct {
	ActionGroupsAPIs   *actiongroupsapis.ActionGroupsAPIsClient
	MetricAlerts       *metricalerts.MetricAlertsClient
	MetricAlertsStatus *metricalertsstatus.MetricAlertsStatusClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	actionGroupsAPIsClient := actiongroupsapis.NewActionGroupsAPIsClientWithBaseURI(endpoint)
	configureAuthFunc(&actionGroupsAPIsClient.Client)

	metricAlertsClient := metricalerts.NewMetricAlertsClientWithBaseURI(endpoint)
	configureAuthFunc(&metricAlertsClient.Client)

	metricAlertsStatusClient := metricalertsstatus.NewMetricAlertsStatusClientWithBaseURI(endpoint)
	configureAuthFunc(&metricAlertsStatusClient.Client)

	return Client{
		ActionGroupsAPIs:   &actionGroupsAPIsClient,
		MetricAlerts:       &metricAlertsClient,
		MetricAlertsStatus: &metricAlertsStatusClient,
	}
}
