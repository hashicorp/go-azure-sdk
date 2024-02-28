package workspacesettings

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceSettingsClient struct {
	Client *resourcemanager.Client
}

func NewWorkspaceSettingsClientWithBaseURI(sdkApi sdkEnv.Api) (*WorkspaceSettingsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "workspacesettings", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating WorkspaceSettingsClient: %+v", err)
	}

	return &WorkspaceSettingsClient{
		Client: client,
	}, nil
}
