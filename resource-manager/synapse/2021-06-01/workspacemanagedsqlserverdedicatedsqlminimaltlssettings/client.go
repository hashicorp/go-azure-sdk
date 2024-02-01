package workspacemanagedsqlserverdedicatedsqlminimaltlssettings

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceManagedSqlServerDedicatedSQLminimalTlsSettingsClient struct {
	Client *resourcemanager.Client
}

func NewWorkspaceManagedSqlServerDedicatedSQLminimalTlsSettingsClientWithBaseURI(sdkApi sdkEnv.Api) (*WorkspaceManagedSqlServerDedicatedSQLminimalTlsSettingsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "workspacemanagedsqlserverdedicatedsqlminimaltlssettings", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating WorkspaceManagedSqlServerDedicatedSQLminimalTlsSettingsClient: %+v", err)
	}

	return &WorkspaceManagedSqlServerDedicatedSQLminimalTlsSettingsClient{
		Client: client,
	}, nil
}
