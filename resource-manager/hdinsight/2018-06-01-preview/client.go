package v2018_06_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/hdinsight/2018-06-01-preview/applications"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hdinsight/2018-06-01-preview/clusters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hdinsight/2018-06-01-preview/configurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hdinsight/2018-06-01-preview/extensions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hdinsight/2018-06-01-preview/promote"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hdinsight/2018-06-01-preview/regions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hdinsight/2018-06-01-preview/scriptactions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hdinsight/2018-06-01-preview/scriptexecutionhistory"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hdinsight/2018-06-01-preview/virtualmachines"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Applications           *applications.ApplicationsClient
	Clusters               *clusters.ClustersClient
	Configurations         *configurations.ConfigurationsClient
	Extensions             *extensions.ExtensionsClient
	Promote                *promote.PromoteClient
	Regions                *regions.RegionsClient
	ScriptActions          *scriptactions.ScriptActionsClient
	ScriptExecutionHistory *scriptexecutionhistory.ScriptExecutionHistoryClient
	VirtualMachines        *virtualmachines.VirtualMachinesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	applicationsClient, err := applications.NewApplicationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Applications client: %+v", err)
	}
	configureFunc(applicationsClient.Client)

	clustersClient, err := clusters.NewClustersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Clusters client: %+v", err)
	}
	configureFunc(clustersClient.Client)

	configurationsClient, err := configurations.NewConfigurationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Configurations client: %+v", err)
	}
	configureFunc(configurationsClient.Client)

	extensionsClient, err := extensions.NewExtensionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Extensions client: %+v", err)
	}
	configureFunc(extensionsClient.Client)

	promoteClient, err := promote.NewPromoteClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Promote client: %+v", err)
	}
	configureFunc(promoteClient.Client)

	regionsClient, err := regions.NewRegionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Regions client: %+v", err)
	}
	configureFunc(regionsClient.Client)

	scriptActionsClient, err := scriptactions.NewScriptActionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ScriptActions client: %+v", err)
	}
	configureFunc(scriptActionsClient.Client)

	scriptExecutionHistoryClient, err := scriptexecutionhistory.NewScriptExecutionHistoryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ScriptExecutionHistory client: %+v", err)
	}
	configureFunc(scriptExecutionHistoryClient.Client)

	virtualMachinesClient, err := virtualmachines.NewVirtualMachinesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualMachines client: %+v", err)
	}
	configureFunc(virtualMachinesClient.Client)

	return &Client{
		Applications:           applicationsClient,
		Clusters:               clustersClient,
		Configurations:         configurationsClient,
		Extensions:             extensionsClient,
		Promote:                promoteClient,
		Regions:                regionsClient,
		ScriptActions:          scriptActionsClient,
		ScriptExecutionHistory: scriptExecutionHistoryClient,
		VirtualMachines:        virtualMachinesClient,
	}, nil
}
