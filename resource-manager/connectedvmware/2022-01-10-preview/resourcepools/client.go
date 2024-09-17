package resourcepools

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourcePoolsClient struct {
	Client *resourcemanager.Client
}

func NewResourcePoolsClientWithBaseURI(sdkApi sdkEnv.Api) (*ResourcePoolsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "resourcepools", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ResourcePoolsClient: %+v", err)
	}

	return &ResourcePoolsClient{
		Client: client,
	}, nil
}
