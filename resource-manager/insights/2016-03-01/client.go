package v2016_03_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2016-03-01/alertruleincidents"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2016-03-01/alertrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2016-03-01/alertrulesapis"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2016-03-01/logprofiles"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2016-03-01/logprofilesapis"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2016-03-01/metricdefinitions"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AlertRuleIncidents *alertruleincidents.AlertRuleIncidentsClient
	AlertRules         *alertrules.AlertRulesClient
	AlertRulesAPIs     *alertrulesapis.AlertRulesAPIsClient
	LogProfiles        *logprofiles.LogProfilesClient
	LogProfilesAPIs    *logprofilesapis.LogProfilesAPIsClient
	MetricDefinitions  *metricdefinitions.MetricDefinitionsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	alertRuleIncidentsClient, err := alertruleincidents.NewAlertRuleIncidentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AlertRuleIncidents client: %+v", err)
	}
	configureFunc(alertRuleIncidentsClient.Client)

	alertRulesAPIsClient, err := alertrulesapis.NewAlertRulesAPIsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AlertRulesAPIs client: %+v", err)
	}
	configureFunc(alertRulesAPIsClient.Client)

	alertRulesClient, err := alertrules.NewAlertRulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AlertRules client: %+v", err)
	}
	configureFunc(alertRulesClient.Client)

	logProfilesAPIsClient, err := logprofilesapis.NewLogProfilesAPIsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LogProfilesAPIs client: %+v", err)
	}
	configureFunc(logProfilesAPIsClient.Client)

	logProfilesClient, err := logprofiles.NewLogProfilesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LogProfiles client: %+v", err)
	}
	configureFunc(logProfilesClient.Client)

	metricDefinitionsClient, err := metricdefinitions.NewMetricDefinitionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MetricDefinitions client: %+v", err)
	}
	configureFunc(metricDefinitionsClient.Client)

	return &Client{
		AlertRuleIncidents: alertRuleIncidentsClient,
		AlertRules:         alertRulesClient,
		AlertRulesAPIs:     alertRulesAPIsClient,
		LogProfiles:        logProfilesClient,
		LogProfilesAPIs:    logProfilesAPIsClient,
		MetricDefinitions:  metricDefinitionsClient,
	}, nil
}
