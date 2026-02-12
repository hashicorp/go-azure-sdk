package v2023_03_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/alertsmanagement/2023-03-01/prometheusrulegroupresources"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	PrometheusRuleGroupResources *prometheusrulegroupresources.PrometheusRuleGroupResourcesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	prometheusRuleGroupResourcesClient, err := prometheusrulegroupresources.NewPrometheusRuleGroupResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrometheusRuleGroupResources client: %+v", err)
	}
	configureFunc(prometheusRuleGroupResourcesClient.Client)

	return &Client{
		PrometheusRuleGroupResources: prometheusRuleGroupResourcesClient,
	}, nil
}
