package connectedresources

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectedResourcesClient struct {
	Client *resourcemanager.Client
}

func NewConnectedResourcesClientWithBaseURI(sdkApi sdkEnv.Api) (*ConnectedResourcesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "connectedresources", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ConnectedResourcesClient: %+v", err)
	}

	return &ConnectedResourcesClient{
		Client: client,
	}, nil
}
