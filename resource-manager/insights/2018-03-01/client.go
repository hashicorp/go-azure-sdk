package v2018_03_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2018-03-01/actiongroupsapis"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2018-03-01/metricalerts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2018-03-01/metricalertsstatus"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	ActionGroupsAPIs   *actiongroupsapis.ActionGroupsAPIsClient
	MetricAlerts       *metricalerts.MetricAlertsClient
	MetricAlertsStatus *metricalertsstatus.MetricAlertsStatusClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	actionGroupsAPIsClient, err := actiongroupsapis.NewActionGroupsAPIsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ActionGroupsAPIs client: %+v", err)
	}
	configureFunc(actionGroupsAPIsClient.Client)

	metricAlertsClient, err := metricalerts.NewMetricAlertsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MetricAlerts client: %+v", err)
	}
	configureFunc(metricAlertsClient.Client)

	metricAlertsStatusClient, err := metricalertsstatus.NewMetricAlertsStatusClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MetricAlertsStatus client: %+v", err)
	}
	configureFunc(metricAlertsStatusClient.Client)

	return &Client{
		ActionGroupsAPIs:   actionGroupsAPIsClient,
		MetricAlerts:       metricAlertsClient,
		MetricAlertsStatus: metricAlertsStatusClient,
	}, nil
}
