package v2022_01_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/servicefabricmanagedcluster/2022-01-01/application"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicefabricmanagedcluster/2022-01-01/applicationtype"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicefabricmanagedcluster/2022-01-01/applicationtypeversion"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicefabricmanagedcluster/2022-01-01/managedcluster"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicefabricmanagedcluster/2022-01-01/managedclusterversion"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicefabricmanagedcluster/2022-01-01/managedvmsizes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicefabricmanagedcluster/2022-01-01/nodetype"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicefabricmanagedcluster/2022-01-01/service"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicefabricmanagedcluster/2022-01-01/services"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Application            *application.ApplicationClient
	ApplicationType        *applicationtype.ApplicationTypeClient
	ApplicationTypeVersion *applicationtypeversion.ApplicationTypeVersionClient
	ManagedCluster         *managedcluster.ManagedClusterClient
	ManagedClusterVersion  *managedclusterversion.ManagedClusterVersionClient
	ManagedVMSizes         *managedvmsizes.ManagedVMSizesClient
	NodeType               *nodetype.NodeTypeClient
	Service                *service.ServiceClient
	Services               *services.ServicesClient
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

	managedClusterClient, err := managedcluster.NewManagedClusterClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedCluster client: %+v", err)
	}
	configureFunc(managedClusterClient.Client)

	managedClusterVersionClient, err := managedclusterversion.NewManagedClusterVersionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedClusterVersion client: %+v", err)
	}
	configureFunc(managedClusterVersionClient.Client)

	managedVMSizesClient, err := managedvmsizes.NewManagedVMSizesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedVMSizes client: %+v", err)
	}
	configureFunc(managedVMSizesClient.Client)

	nodeTypeClient, err := nodetype.NewNodeTypeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NodeType client: %+v", err)
	}
	configureFunc(nodeTypeClient.Client)

	serviceClient, err := service.NewServiceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Service client: %+v", err)
	}
	configureFunc(serviceClient.Client)

	servicesClient, err := services.NewServicesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Services client: %+v", err)
	}
	configureFunc(servicesClient.Client)

	return &Client{
		Application:            applicationClient,
		ApplicationType:        applicationTypeClient,
		ApplicationTypeVersion: applicationTypeVersionClient,
		ManagedCluster:         managedClusterClient,
		ManagedClusterVersion:  managedClusterVersionClient,
		ManagedVMSizes:         managedVMSizesClient,
		NodeType:               nodeTypeClient,
		Service:                serviceClient,
		Services:               servicesClient,
	}, nil
}
