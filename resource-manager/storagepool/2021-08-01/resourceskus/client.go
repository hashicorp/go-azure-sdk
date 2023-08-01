package resourceskus

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceSkusClient struct {
	Client *resourcemanager.Client
}

func NewResourceSkusClientWithBaseURI(sdkApi sdkEnv.Api) (*ResourceSkusClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "resourceskus", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ResourceSkusClient: %+v", err)
	}

	return &ResourceSkusClient{
		Client: client,
	}, nil
}
