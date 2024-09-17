package workspacemanagedsqlserver

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceManagedSqlServerClient struct {
	Client *resourcemanager.Client
}

func NewWorkspaceManagedSqlServerClientWithBaseURI(sdkApi sdkEnv.Api) (*WorkspaceManagedSqlServerClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "workspacemanagedsqlserver", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating WorkspaceManagedSqlServerClient: %+v", err)
	}

	return &WorkspaceManagedSqlServerClient{
		Client: client,
	}, nil
}
