package datacontainerregistry

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataContainerRegistryClient struct {
	Client *resourcemanager.Client
}

func NewDataContainerRegistryClientWithBaseURI(sdkApi sdkEnv.Api) (*DataContainerRegistryClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "datacontainerregistry", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DataContainerRegistryClient: %+v", err)
	}

	return &DataContainerRegistryClient{
		Client: client,
	}, nil
}
