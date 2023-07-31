package v2023_06_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2023-06-01/apikey"
	"github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2023-06-01/deploymentinfo"
	"github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2023-06-01/deploymentupdate"
	"github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2023-06-01/elasticversions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2023-06-01/monitoredresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2023-06-01/monitorsresource"
	"github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2023-06-01/monitorupgradableversions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2023-06-01/rules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2023-06-01/trafficfilter"
	"github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2023-06-01/vmcollectionupdate"
	"github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2023-06-01/vmhhostlist"
	"github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2023-06-01/vmingestiondetails"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	ApiKey                    *apikey.ApiKeyClient
	DeploymentInfo            *deploymentinfo.DeploymentInfoClient
	DeploymentUpdate          *deploymentupdate.DeploymentUpdateClient
	ElasticVersions           *elasticversions.ElasticVersionsClient
	MonitorUpgradableVersions *monitorupgradableversions.MonitorUpgradableVersionsClient
	MonitoredResources        *monitoredresources.MonitoredResourcesClient
	MonitorsResource          *monitorsresource.MonitorsResourceClient
	Rules                     *rules.RulesClient
	TrafficFilter             *trafficfilter.TrafficFilterClient
	VMCollectionUpdate        *vmcollectionupdate.VMCollectionUpdateClient
	VMHHostList               *vmhhostlist.VMHHostListClient
	VMIngestionDetails        *vmingestiondetails.VMIngestionDetailsClient
}

func NewClientWithBaseURI(api sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	apiKeyClient, err := apikey.NewApiKeyClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ApiKey client: %+v", err)
	}
	configureFunc(apiKeyClient.Client)

	deploymentInfoClient, err := deploymentinfo.NewDeploymentInfoClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building DeploymentInfo client: %+v", err)
	}
	configureFunc(deploymentInfoClient.Client)

	deploymentUpdateClient, err := deploymentupdate.NewDeploymentUpdateClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building DeploymentUpdate client: %+v", err)
	}
	configureFunc(deploymentUpdateClient.Client)

	elasticVersionsClient, err := elasticversions.NewElasticVersionsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ElasticVersions client: %+v", err)
	}
	configureFunc(elasticVersionsClient.Client)

	monitorUpgradableVersionsClient, err := monitorupgradableversions.NewMonitorUpgradableVersionsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building MonitorUpgradableVersions client: %+v", err)
	}
	configureFunc(monitorUpgradableVersionsClient.Client)

	monitoredResourcesClient, err := monitoredresources.NewMonitoredResourcesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building MonitoredResources client: %+v", err)
	}
	configureFunc(monitoredResourcesClient.Client)

	monitorsResourceClient, err := monitorsresource.NewMonitorsResourceClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building MonitorsResource client: %+v", err)
	}
	configureFunc(monitorsResourceClient.Client)

	rulesClient, err := rules.NewRulesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Rules client: %+v", err)
	}
	configureFunc(rulesClient.Client)

	trafficFilterClient, err := trafficfilter.NewTrafficFilterClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building TrafficFilter client: %+v", err)
	}
	configureFunc(trafficFilterClient.Client)

	vMCollectionUpdateClient, err := vmcollectionupdate.NewVMCollectionUpdateClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building VMCollectionUpdate client: %+v", err)
	}
	configureFunc(vMCollectionUpdateClient.Client)

	vMHHostListClient, err := vmhhostlist.NewVMHHostListClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building VMHHostList client: %+v", err)
	}
	configureFunc(vMHHostListClient.Client)

	vMIngestionDetailsClient, err := vmingestiondetails.NewVMIngestionDetailsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building VMIngestionDetails client: %+v", err)
	}
	configureFunc(vMIngestionDetailsClient.Client)

	return &Client{
		ApiKey:                    apiKeyClient,
		DeploymentInfo:            deploymentInfoClient,
		DeploymentUpdate:          deploymentUpdateClient,
		ElasticVersions:           elasticVersionsClient,
		MonitorUpgradableVersions: monitorUpgradableVersionsClient,
		MonitoredResources:        monitoredResourcesClient,
		MonitorsResource:          monitorsResourceClient,
		Rules:                     rulesClient,
		TrafficFilter:             trafficFilterClient,
		VMCollectionUpdate:        vMCollectionUpdateClient,
		VMHHostList:               vMHHostListClient,
		VMIngestionDetails:        vMIngestionDetailsClient,
	}, nil
}
