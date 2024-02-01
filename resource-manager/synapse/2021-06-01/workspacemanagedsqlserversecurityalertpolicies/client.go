package workspacemanagedsqlserversecurityalertpolicies

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceManagedSqlServerSecurityAlertPoliciesClient struct {
	Client *resourcemanager.Client
}

func NewWorkspaceManagedSqlServerSecurityAlertPoliciesClientWithBaseURI(sdkApi sdkEnv.Api) (*WorkspaceManagedSqlServerSecurityAlertPoliciesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "workspacemanagedsqlserversecurityalertpolicies", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating WorkspaceManagedSqlServerSecurityAlertPoliciesClient: %+v", err)
	}

	return &WorkspaceManagedSqlServerSecurityAlertPoliciesClient{
		Client: client,
	}, nil
}
