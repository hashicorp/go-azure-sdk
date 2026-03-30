package clustermetricsconfigurations

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClusterMetricsConfigurationsClient struct {
	Client *resourcemanager.Client
}

func NewClusterMetricsConfigurationsClientWithBaseURI(sdkApi sdkEnv.Api) (*ClusterMetricsConfigurationsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "clustermetricsconfigurations", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ClusterMetricsConfigurationsClient: %+v", err)
	}

	return &ClusterMetricsConfigurationsClient{
		Client: client,
	}, nil
}
