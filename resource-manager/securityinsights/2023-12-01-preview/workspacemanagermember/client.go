package workspacemanagermember

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceManagerMemberClient struct {
	Client *resourcemanager.Client
}

func NewWorkspaceManagerMemberClientWithBaseURI(sdkApi sdkEnv.Api) (*WorkspaceManagerMemberClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "workspacemanagermember", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating WorkspaceManagerMemberClient: %+v", err)
	}

	return &WorkspaceManagerMemberClient{
		Client: client,
	}, nil
}
