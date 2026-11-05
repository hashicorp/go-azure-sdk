package sandboxgroups

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SandboxGroupsClient struct {
	Client *resourcemanager.Client
}

func NewSandboxGroupsClientWithBaseURI(sdkApi sdkEnv.Api) (*SandboxGroupsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "sandboxgroups", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SandboxGroupsClient: %+v", err)
	}

	return &SandboxGroupsClient{
		Client: client,
	}, nil
}
