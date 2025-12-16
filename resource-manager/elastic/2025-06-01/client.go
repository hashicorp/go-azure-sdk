package v2025_06_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2025-06-01/elasticmonitorresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2025-06-01/elastics"
	"github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2025-06-01/monitoredsubscriptions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2025-06-01/openaiintegrationrpmodels"
	"github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2025-06-01/tagrules"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	ElasticMonitorResources   *elasticmonitorresources.ElasticMonitorResourcesClient
	Elastics                  *elastics.ElasticsClient
	MonitoredSubscriptions    *monitoredsubscriptions.MonitoredSubscriptionsClient
	OpenAIIntegrationRPModels *openaiintegrationrpmodels.OpenAIIntegrationRPModelsClient
	TagRules                  *tagrules.TagRulesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	elasticMonitorResourcesClient, err := elasticmonitorresources.NewElasticMonitorResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ElasticMonitorResources client: %+v", err)
	}
	configureFunc(elasticMonitorResourcesClient.Client)

	elasticsClient, err := elastics.NewElasticsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Elastics client: %+v", err)
	}
	configureFunc(elasticsClient.Client)

	monitoredSubscriptionsClient, err := monitoredsubscriptions.NewMonitoredSubscriptionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MonitoredSubscriptions client: %+v", err)
	}
	configureFunc(monitoredSubscriptionsClient.Client)

	openAIIntegrationRPModelsClient, err := openaiintegrationrpmodels.NewOpenAIIntegrationRPModelsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OpenAIIntegrationRPModels client: %+v", err)
	}
	configureFunc(openAIIntegrationRPModelsClient.Client)

	tagRulesClient, err := tagrules.NewTagRulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TagRules client: %+v", err)
	}
	configureFunc(tagRulesClient.Client)

	return &Client{
		ElasticMonitorResources:   elasticMonitorResourcesClient,
		Elastics:                  elasticsClient,
		MonitoredSubscriptions:    monitoredSubscriptionsClient,
		OpenAIIntegrationRPModels: openAIIntegrationRPModelsClient,
		TagRules:                  tagRulesClient,
	}, nil
}
