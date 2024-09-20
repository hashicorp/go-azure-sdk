package workspacemanagedsqlserverblobauditing

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceManagedSqlServerBlobAuditingClient struct {
	Client *resourcemanager.Client
}

func NewWorkspaceManagedSqlServerBlobAuditingClientWithBaseURI(sdkApi sdkEnv.Api) (*WorkspaceManagedSqlServerBlobAuditingClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "workspacemanagedsqlserverblobauditing", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating WorkspaceManagedSqlServerBlobAuditingClient: %+v", err)
	}

	return &WorkspaceManagedSqlServerBlobAuditingClient{
		Client: client,
	}, nil
}
