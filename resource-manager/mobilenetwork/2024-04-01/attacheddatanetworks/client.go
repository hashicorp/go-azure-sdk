package attacheddatanetworks

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttachedDataNetworksClient struct {
	Client *resourcemanager.Client
}

func NewAttachedDataNetworksClientWithBaseURI(sdkApi sdkEnv.Api) (*AttachedDataNetworksClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "attacheddatanetworks", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AttachedDataNetworksClient: %+v", err)
	}

	return &AttachedDataNetworksClient{
		Client: client,
	}, nil
}
