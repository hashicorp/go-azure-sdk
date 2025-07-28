package v2025_03_31

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/hardwaresecuritymodules/2025-03-31/cloudhsmclusters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hardwaresecuritymodules/2025-03-31/dedicatedhsms"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hardwaresecuritymodules/2025-03-31/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	CloudHsmClusters           *cloudhsmclusters.CloudHsmClustersClient
	DedicatedHsms              *dedicatedhsms.DedicatedHsmsClient
	PrivateEndpointConnections *privateendpointconnections.PrivateEndpointConnectionsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	cloudHsmClustersClient, err := cloudhsmclusters.NewCloudHsmClustersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CloudHsmClusters client: %+v", err)
	}
	configureFunc(cloudHsmClustersClient.Client)

	dedicatedHsmsClient, err := dedicatedhsms.NewDedicatedHsmsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DedicatedHsms client: %+v", err)
	}
	configureFunc(dedicatedHsmsClient.Client)

	privateEndpointConnectionsClient, err := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateEndpointConnections client: %+v", err)
	}
	configureFunc(privateEndpointConnectionsClient.Client)

	return &Client{
		CloudHsmClusters:           cloudHsmClustersClient,
		DedicatedHsms:              dedicatedHsmsClient,
		PrivateEndpointConnections: privateEndpointConnectionsClient,
	}, nil
}
