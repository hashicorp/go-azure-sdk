package v2workspaceconnectionresource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type V2WorkspaceConnectionResourceClient struct {
	Client *resourcemanager.Client
}

func NewV2WorkspaceConnectionResourceClientWithBaseURI(sdkApi sdkEnv.Api) (*V2WorkspaceConnectionResourceClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "v2workspaceconnectionresource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating V2WorkspaceConnectionResourceClient: %+v", err)
	}

	return &V2WorkspaceConnectionResourceClient{
		Client: client,
	}, nil
}
