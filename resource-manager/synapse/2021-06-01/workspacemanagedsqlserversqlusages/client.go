package workspacemanagedsqlserversqlusages

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceManagedSqlServerSqlUsagesClient struct {
	Client *resourcemanager.Client
}

func NewWorkspaceManagedSqlServerSqlUsagesClientWithBaseURI(sdkApi sdkEnv.Api) (*WorkspaceManagedSqlServerSqlUsagesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "workspacemanagedsqlserversqlusages", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating WorkspaceManagedSqlServerSqlUsagesClient: %+v", err)
	}

	return &WorkspaceManagedSqlServerSqlUsagesClient{
		Client: client,
	}, nil
}
