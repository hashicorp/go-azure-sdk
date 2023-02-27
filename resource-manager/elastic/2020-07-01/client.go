// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2020_07_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2020-07-01/deploymentinfo"
	"github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2020-07-01/monitoredresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2020-07-01/monitorsresource"
	"github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2020-07-01/rules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2020-07-01/vmcollectionupdate"
	"github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2020-07-01/vmhhostlist"
	"github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2020-07-01/vmingestiondetails"
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

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	deploymentInfoClient := deploymentinfo.NewDeploymentInfoClientWithBaseURI(endpoint)
	configureAuthFunc(&deploymentInfoClient.Client)

	monitoredResourcesClient := monitoredresources.NewMonitoredResourcesClientWithBaseURI(endpoint)
	configureAuthFunc(&monitoredResourcesClient.Client)

	monitorsResourceClient := monitorsresource.NewMonitorsResourceClientWithBaseURI(endpoint)
	configureAuthFunc(&monitorsResourceClient.Client)

	rulesClient := rules.NewRulesClientWithBaseURI(endpoint)
	configureAuthFunc(&rulesClient.Client)

	vMCollectionUpdateClient := vmcollectionupdate.NewVMCollectionUpdateClientWithBaseURI(endpoint)
	configureAuthFunc(&vMCollectionUpdateClient.Client)

	vMHHostListClient := vmhhostlist.NewVMHHostListClientWithBaseURI(endpoint)
	configureAuthFunc(&vMHHostListClient.Client)

	vMIngestionDetailsClient := vmingestiondetails.NewVMIngestionDetailsClientWithBaseURI(endpoint)
	configureAuthFunc(&vMIngestionDetailsClient.Client)

	return Client{
		DeploymentInfo:     &deploymentInfoClient,
		MonitoredResources: &monitoredResourcesClient,
		MonitorsResource:   &monitorsResourceClient,
		Rules:              &rulesClient,
		VMCollectionUpdate: &vMCollectionUpdateClient,
		VMHHostList:        &vMHHostListClient,
		VMIngestionDetails: &vMIngestionDetailsClient,
	}
}
