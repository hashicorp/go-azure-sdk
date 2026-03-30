package v2025_09_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/networkcloud/2025-09-01/agentpools"
	"github.com/hashicorp/go-azure-sdk/resource-manager/networkcloud/2025-09-01/baremetalmachinekeysets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/networkcloud/2025-09-01/baremetalmachines"
	"github.com/hashicorp/go-azure-sdk/resource-manager/networkcloud/2025-09-01/bmckeysets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/networkcloud/2025-09-01/cloudservicesnetworks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/networkcloud/2025-09-01/clustermanagers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/networkcloud/2025-09-01/clustermetricsconfigurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/networkcloud/2025-09-01/clusters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/networkcloud/2025-09-01/consoles"
	"github.com/hashicorp/go-azure-sdk/resource-manager/networkcloud/2025-09-01/kubernetesclusterfeatures"
	"github.com/hashicorp/go-azure-sdk/resource-manager/networkcloud/2025-09-01/kubernetesclusters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/networkcloud/2025-09-01/l2networks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/networkcloud/2025-09-01/l3networks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/networkcloud/2025-09-01/racks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/networkcloud/2025-09-01/rackskus"
	"github.com/hashicorp/go-azure-sdk/resource-manager/networkcloud/2025-09-01/storageappliances"
	"github.com/hashicorp/go-azure-sdk/resource-manager/networkcloud/2025-09-01/trunkednetworks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/networkcloud/2025-09-01/virtualmachines"
	"github.com/hashicorp/go-azure-sdk/resource-manager/networkcloud/2025-09-01/volumes"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AgentPools                   *agentpools.AgentPoolsClient
	BareMetalMachineKeySets      *baremetalmachinekeysets.BareMetalMachineKeySetsClient
	BareMetalMachines            *baremetalmachines.BareMetalMachinesClient
	BmcKeySets                   *bmckeysets.BmcKeySetsClient
	CloudServicesNetworks        *cloudservicesnetworks.CloudServicesNetworksClient
	ClusterManagers              *clustermanagers.ClusterManagersClient
	ClusterMetricsConfigurations *clustermetricsconfigurations.ClusterMetricsConfigurationsClient
	Clusters                     *clusters.ClustersClient
	Consoles                     *consoles.ConsolesClient
	KubernetesClusterFeatures    *kubernetesclusterfeatures.KubernetesClusterFeaturesClient
	KubernetesClusters           *kubernetesclusters.KubernetesClustersClient
	L2Networks                   *l2networks.L2NetworksClient
	L3Networks                   *l3networks.L3NetworksClient
	RackSkus                     *rackskus.RackSkusClient
	Racks                        *racks.RacksClient
	StorageAppliances            *storageappliances.StorageAppliancesClient
	TrunkedNetworks              *trunkednetworks.TrunkedNetworksClient
	VirtualMachines              *virtualmachines.VirtualMachinesClient
	Volumes                      *volumes.VolumesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	agentPoolsClient, err := agentpools.NewAgentPoolsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AgentPools client: %+v", err)
	}
	configureFunc(agentPoolsClient.Client)

	bareMetalMachineKeySetsClient, err := baremetalmachinekeysets.NewBareMetalMachineKeySetsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BareMetalMachineKeySets client: %+v", err)
	}
	configureFunc(bareMetalMachineKeySetsClient.Client)

	bareMetalMachinesClient, err := baremetalmachines.NewBareMetalMachinesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BareMetalMachines client: %+v", err)
	}
	configureFunc(bareMetalMachinesClient.Client)

	bmcKeySetsClient, err := bmckeysets.NewBmcKeySetsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BmcKeySets client: %+v", err)
	}
	configureFunc(bmcKeySetsClient.Client)

	cloudServicesNetworksClient, err := cloudservicesnetworks.NewCloudServicesNetworksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CloudServicesNetworks client: %+v", err)
	}
	configureFunc(cloudServicesNetworksClient.Client)

	clusterManagersClient, err := clustermanagers.NewClusterManagersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ClusterManagers client: %+v", err)
	}
	configureFunc(clusterManagersClient.Client)

	clusterMetricsConfigurationsClient, err := clustermetricsconfigurations.NewClusterMetricsConfigurationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ClusterMetricsConfigurations client: %+v", err)
	}
	configureFunc(clusterMetricsConfigurationsClient.Client)

	clustersClient, err := clusters.NewClustersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Clusters client: %+v", err)
	}
	configureFunc(clustersClient.Client)

	consolesClient, err := consoles.NewConsolesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Consoles client: %+v", err)
	}
	configureFunc(consolesClient.Client)

	kubernetesClusterFeaturesClient, err := kubernetesclusterfeatures.NewKubernetesClusterFeaturesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building KubernetesClusterFeatures client: %+v", err)
	}
	configureFunc(kubernetesClusterFeaturesClient.Client)

	kubernetesClustersClient, err := kubernetesclusters.NewKubernetesClustersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building KubernetesClusters client: %+v", err)
	}
	configureFunc(kubernetesClustersClient.Client)

	l2NetworksClient, err := l2networks.NewL2NetworksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building L2Networks client: %+v", err)
	}
	configureFunc(l2NetworksClient.Client)

	l3NetworksClient, err := l3networks.NewL3NetworksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building L3Networks client: %+v", err)
	}
	configureFunc(l3NetworksClient.Client)

	rackSkusClient, err := rackskus.NewRackSkusClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RackSkus client: %+v", err)
	}
	configureFunc(rackSkusClient.Client)

	racksClient, err := racks.NewRacksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Racks client: %+v", err)
	}
	configureFunc(racksClient.Client)

	storageAppliancesClient, err := storageappliances.NewStorageAppliancesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building StorageAppliances client: %+v", err)
	}
	configureFunc(storageAppliancesClient.Client)

	trunkedNetworksClient, err := trunkednetworks.NewTrunkedNetworksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TrunkedNetworks client: %+v", err)
	}
	configureFunc(trunkedNetworksClient.Client)

	virtualMachinesClient, err := virtualmachines.NewVirtualMachinesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualMachines client: %+v", err)
	}
	configureFunc(virtualMachinesClient.Client)

	volumesClient, err := volumes.NewVolumesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Volumes client: %+v", err)
	}
	configureFunc(volumesClient.Client)

	return &Client{
		AgentPools:                   agentPoolsClient,
		BareMetalMachineKeySets:      bareMetalMachineKeySetsClient,
		BareMetalMachines:            bareMetalMachinesClient,
		BmcKeySets:                   bmcKeySetsClient,
		CloudServicesNetworks:        cloudServicesNetworksClient,
		ClusterManagers:              clusterManagersClient,
		ClusterMetricsConfigurations: clusterMetricsConfigurationsClient,
		Clusters:                     clustersClient,
		Consoles:                     consolesClient,
		KubernetesClusterFeatures:    kubernetesClusterFeaturesClient,
		KubernetesClusters:           kubernetesClustersClient,
		L2Networks:                   l2NetworksClient,
		L3Networks:                   l3NetworksClient,
		RackSkus:                     rackSkusClient,
		Racks:                        racksClient,
		StorageAppliances:            storageAppliancesClient,
		TrunkedNetworks:              trunkedNetworksClient,
		VirtualMachines:              virtualMachinesClient,
		Volumes:                      volumesClient,
	}, nil
}
