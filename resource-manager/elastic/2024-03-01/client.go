package v2024_03_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2024-03-01/apikey"
	"github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2024-03-01/connectedresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2024-03-01/deploymentinfo"
	"github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2024-03-01/deploymentupdate"
	"github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2024-03-01/elasticversions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2024-03-01/monitoredresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2024-03-01/monitorsresource"
	"github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2024-03-01/monitorupgradableversions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2024-03-01/openaiintegration"
	"github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2024-03-01/rules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2024-03-01/trafficfilter"
	"github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2024-03-01/userorganization"
	"github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2024-03-01/vmcollectionupdate"
	"github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2024-03-01/vmhhostlist"
	"github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2024-03-01/vmingestiondetails"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	ApiKey                    *apikey.ApiKeyClient
	ConnectedResources        *connectedresources.ConnectedResourcesClient
	DeploymentInfo            *deploymentinfo.DeploymentInfoClient
	DeploymentUpdate          *deploymentupdate.DeploymentUpdateClient
	ElasticVersions           *elasticversions.ElasticVersionsClient
	MonitorUpgradableVersions *monitorupgradableversions.MonitorUpgradableVersionsClient
	MonitoredResources        *monitoredresources.MonitoredResourcesClient
	MonitorsResource          *monitorsresource.MonitorsResourceClient
	OpenAIIntegration         *openaiintegration.OpenAIIntegrationClient
	Rules                     *rules.RulesClient
	TrafficFilter             *trafficfilter.TrafficFilterClient
	UserOrganization          *userorganization.UserOrganizationClient
	VMCollectionUpdate        *vmcollectionupdate.VMCollectionUpdateClient
	VMHHostList               *vmhhostlist.VMHHostListClient
	VMIngestionDetails        *vmingestiondetails.VMIngestionDetailsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	apiKeyClient, err := apikey.NewApiKeyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ApiKey client: %+v", err)
	}
	configureFunc(apiKeyClient.Client)

	connectedResourcesClient, err := connectedresources.NewConnectedResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConnectedResources client: %+v", err)
	}
	configureFunc(connectedResourcesClient.Client)

	deploymentInfoClient, err := deploymentinfo.NewDeploymentInfoClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeploymentInfo client: %+v", err)
	}
	configureFunc(deploymentInfoClient.Client)

	deploymentUpdateClient, err := deploymentupdate.NewDeploymentUpdateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeploymentUpdate client: %+v", err)
	}
	configureFunc(deploymentUpdateClient.Client)

	elasticVersionsClient, err := elasticversions.NewElasticVersionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ElasticVersions client: %+v", err)
	}
	configureFunc(elasticVersionsClient.Client)

	monitorUpgradableVersionsClient, err := monitorupgradableversions.NewMonitorUpgradableVersionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MonitorUpgradableVersions client: %+v", err)
	}
	configureFunc(monitorUpgradableVersionsClient.Client)

	monitoredResourcesClient, err := monitoredresources.NewMonitoredResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MonitoredResources client: %+v", err)
	}
	configureFunc(monitoredResourcesClient.Client)

	monitorsResourceClient, err := monitorsresource.NewMonitorsResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MonitorsResource client: %+v", err)
	}
	configureFunc(monitorsResourceClient.Client)

	openAIIntegrationClient, err := openaiintegration.NewOpenAIIntegrationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OpenAIIntegration client: %+v", err)
	}
	configureFunc(openAIIntegrationClient.Client)

	rulesClient, err := rules.NewRulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Rules client: %+v", err)
	}
	configureFunc(rulesClient.Client)

	trafficFilterClient, err := trafficfilter.NewTrafficFilterClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TrafficFilter client: %+v", err)
	}
	configureFunc(trafficFilterClient.Client)

	userOrganizationClient, err := userorganization.NewUserOrganizationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOrganization client: %+v", err)
	}
	configureFunc(userOrganizationClient.Client)

	vMCollectionUpdateClient, err := vmcollectionupdate.NewVMCollectionUpdateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VMCollectionUpdate client: %+v", err)
	}
	configureFunc(vMCollectionUpdateClient.Client)

	vMHHostListClient, err := vmhhostlist.NewVMHHostListClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VMHHostList client: %+v", err)
	}
	configureFunc(vMHHostListClient.Client)

	vMIngestionDetailsClient, err := vmingestiondetails.NewVMIngestionDetailsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VMIngestionDetails client: %+v", err)
	}
	configureFunc(vMIngestionDetailsClient.Client)

	return &Client{
		ApiKey:                    apiKeyClient,
		ConnectedResources:        connectedResourcesClient,
		DeploymentInfo:            deploymentInfoClient,
		DeploymentUpdate:          deploymentUpdateClient,
		ElasticVersions:           elasticVersionsClient,
		MonitorUpgradableVersions: monitorUpgradableVersionsClient,
		MonitoredResources:        monitoredResourcesClient,
		MonitorsResource:          monitorsResourceClient,
		OpenAIIntegration:         openAIIntegrationClient,
		Rules:                     rulesClient,
		TrafficFilter:             trafficFilterClient,
		UserOrganization:          userOrganizationClient,
		VMCollectionUpdate:        vMCollectionUpdateClient,
		VMHHostList:               vMHHostListClient,
		VMIngestionDetails:        vMIngestionDetailsClient,
	}, nil
}
