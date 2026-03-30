package kubernetesclusters

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KubernetesClustersClient struct {
	Client *resourcemanager.Client
}

func NewKubernetesClustersClientWithBaseURI(sdkApi sdkEnv.Api) (*KubernetesClustersClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "kubernetesclusters", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating KubernetesClustersClient: %+v", err)
	}

	return &KubernetesClustersClient{
		Client: client,
	}, nil
}
