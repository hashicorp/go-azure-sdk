package clustermanagers

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClusterManagersClient struct {
	Client *resourcemanager.Client
}

func NewClusterManagersClientWithBaseURI(sdkApi sdkEnv.Api) (*ClusterManagersClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "clustermanagers", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ClusterManagersClient: %+v", err)
	}

	return &ClusterManagersClient{
		Client: client,
	}, nil
}
