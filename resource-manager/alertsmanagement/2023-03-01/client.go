package v2023_03_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/alertsmanagement/2023-03-01/prometheusrulegroups"
)

type Client struct {
	PrometheusRuleGroups *prometheusrulegroups.PrometheusRuleGroupsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	prometheusRuleGroupsClient := prometheusrulegroups.NewPrometheusRuleGroupsClientWithBaseURI(endpoint)
	configureAuthFunc(&prometheusRuleGroupsClient.Client)

	return Client{
		PrometheusRuleGroups: &prometheusRuleGroupsClient,
	}
}
