package v2021_06_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/servicefabric/2021-06-01/application"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicefabric/2021-06-01/applicationtype"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicefabric/2021-06-01/applicationtypeversion"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicefabric/2021-06-01/cluster"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicefabric/2021-06-01/clusterversion"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicefabric/2021-06-01/listupgradableversions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicefabric/2021-06-01/service"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Application            *application.ApplicationClient
	ApplicationType        *applicationtype.ApplicationTypeClient
	ApplicationTypeVersion *applicationtypeversion.ApplicationTypeVersionClient
	Cluster                *cluster.ClusterClient
	ClusterVersion         *clusterversion.ClusterVersionClient
	ListUpgradableVersions *listupgradableversions.ListUpgradableVersionsClient
	Service                *service.ServiceClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	applicationClient, err := application.NewApplicationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Application client: %+v", err)
	}
	configureFunc(applicationClient.Client)

	applicationTypeClient, err := applicationtype.NewApplicationTypeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ApplicationType client: %+v", err)
	}
	configureFunc(applicationTypeClient.Client)

	applicationTypeVersionClient, err := applicationtypeversion.NewApplicationTypeVersionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ApplicationTypeVersion client: %+v", err)
	}
	configureFunc(applicationTypeVersionClient.Client)

	clusterClient, err := cluster.NewClusterClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Cluster client: %+v", err)
	}
	configureFunc(clusterClient.Client)

	clusterVersionClient, err := clusterversion.NewClusterVersionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ClusterVersion client: %+v", err)
	}
	configureFunc(clusterVersionClient.Client)

	listUpgradableVersionsClient, err := listupgradableversions.NewListUpgradableVersionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ListUpgradableVersions client: %+v", err)
	}
	configureFunc(listUpgradableVersionsClient.Client)

	serviceClient, err := service.NewServiceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Service client: %+v", err)
	}
	configureFunc(serviceClient.Client)

	return &Client{
		Application:            applicationClient,
		ApplicationType:        applicationTypeClient,
		ApplicationTypeVersion: applicationTypeVersionClient,
		Cluster:                clusterClient,
		ClusterVersion:         clusterVersionClient,
		ListUpgradableVersions: listUpgradableVersionsClient,
		Service:                serviceClient,
	}, nil
}
