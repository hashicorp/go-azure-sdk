package workspacemanagerconfigurations

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceManagerConfigurationsClient struct {
	Client *resourcemanager.Client
}

func NewWorkspaceManagerConfigurationsClientWithBaseURI(sdkApi sdkEnv.Api) (*WorkspaceManagerConfigurationsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "workspacemanagerconfigurations", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating WorkspaceManagerConfigurationsClient: %+v", err)
	}

	return &WorkspaceManagerConfigurationsClient{
		Client: client,
	}, nil
}
