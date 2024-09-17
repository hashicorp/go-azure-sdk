package synapselinkworkspaces

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynapseLinkWorkspacesClient struct {
	Client *resourcemanager.Client
}

func NewSynapseLinkWorkspacesClientWithBaseURI(sdkApi sdkEnv.Api) (*SynapseLinkWorkspacesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "synapselinkworkspaces", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SynapseLinkWorkspacesClient: %+v", err)
	}

	return &SynapseLinkWorkspacesClient{
		Client: client,
	}, nil
}
