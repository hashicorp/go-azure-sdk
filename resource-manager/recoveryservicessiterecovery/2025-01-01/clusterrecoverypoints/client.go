package clusterrecoverypoints

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClusterRecoveryPointsClient struct {
	Client *resourcemanager.Client
}

func NewClusterRecoveryPointsClientWithBaseURI(sdkApi sdkEnv.Api) (*ClusterRecoveryPointsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "clusterrecoverypoints", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ClusterRecoveryPointsClient: %+v", err)
	}

	return &ClusterRecoveryPointsClient{
		Client: client,
	}, nil
}
