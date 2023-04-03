package v2016_03_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2016-03-01/alertruleincidents"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2016-03-01/alertrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2016-03-01/alertrulesapis"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2016-03-01/logprofiles"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2016-03-01/logprofilesapis"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2016-03-01/metricdefinitions"
)

type Client struct {
	AlertRuleIncidents *alertruleincidents.AlertRuleIncidentsClient
	AlertRules         *alertrules.AlertRulesClient
	AlertRulesAPIs     *alertrulesapis.AlertRulesAPIsClient
	LogProfiles        *logprofiles.LogProfilesClient
	LogProfilesAPIs    *logprofilesapis.LogProfilesAPIsClient
	MetricDefinitions  *metricdefinitions.MetricDefinitionsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	alertRuleIncidentsClient := alertruleincidents.NewAlertRuleIncidentsClientWithBaseURI(endpoint)
	configureAuthFunc(&alertRuleIncidentsClient.Client)

	alertRulesAPIsClient := alertrulesapis.NewAlertRulesAPIsClientWithBaseURI(endpoint)
	configureAuthFunc(&alertRulesAPIsClient.Client)

	alertRulesClient := alertrules.NewAlertRulesClientWithBaseURI(endpoint)
	configureAuthFunc(&alertRulesClient.Client)

	logProfilesAPIsClient := logprofilesapis.NewLogProfilesAPIsClientWithBaseURI(endpoint)
	configureAuthFunc(&logProfilesAPIsClient.Client)

	logProfilesClient := logprofiles.NewLogProfilesClientWithBaseURI(endpoint)
	configureAuthFunc(&logProfilesClient.Client)

	metricDefinitionsClient := metricdefinitions.NewMetricDefinitionsClientWithBaseURI(endpoint)
	configureAuthFunc(&metricDefinitionsClient.Client)

	return Client{
		AlertRuleIncidents: &alertRuleIncidentsClient,
		AlertRules:         &alertRulesClient,
		AlertRulesAPIs:     &alertRulesAPIsClient,
		LogProfiles:        &logProfilesClient,
		LogProfilesAPIs:    &logProfilesAPIsClient,
		MetricDefinitions:  &metricDefinitionsClient,
	}
}
