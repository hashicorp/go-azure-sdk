package v2020_07_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2020-07-01/deploymentinfo"
	"github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2020-07-01/monitoredresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2020-07-01/monitorsresource"
	"github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2020-07-01/rules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2020-07-01/vmcollectionupdate"
	"github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2020-07-01/vmhhostlist"
	"github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2020-07-01/vmingestiondetails"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	DeploymentInfo     *deploymentinfo.DeploymentInfoClient
	MonitoredResources *monitoredresources.MonitoredResourcesClient
	MonitorsResource   *monitorsresource.MonitorsResourceClient
	Rules              *rules.RulesClient
	VMCollectionUpdate *vmcollectionupdate.VMCollectionUpdateClient
	VMHHostList        *vmhhostlist.VMHHostListClient
	VMIngestionDetails *vmingestiondetails.VMIngestionDetailsClient
}

func NewClientWithBaseURI(api environments.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	deploymentInfoClient, err := deploymentinfo.NewDeploymentInfoClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building DeploymentInfo client: %+v", err)
	}
	configureFunc(deploymentInfoClient.Client)

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
		DeploymentInfo:     deploymentInfoClient,
		MonitoredResources: monitoredResourcesClient,
		MonitorsResource:   monitorsResourceClient,
		Rules:              rulesClient,
		VMCollectionUpdate: vMCollectionUpdateClient,
		VMHHostList:        vMHHostListClient,
		VMIngestionDetails: vMIngestionDetailsClient,
	}, nil
}
