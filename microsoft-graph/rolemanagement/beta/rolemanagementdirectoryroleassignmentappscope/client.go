package rolemanagementdirectoryroleassignmentappscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementDirectoryRoleAssignmentAppScopeClient struct {
	Client *msgraph.Client
}

func NewRoleManagementDirectoryRoleAssignmentAppScopeClientWithBaseURI(api sdkEnv.Api) (*RoleManagementDirectoryRoleAssignmentAppScopeClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagementdirectoryroleassignmentappscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementDirectoryRoleAssignmentAppScopeClient: %+v", err)
	}

	return &RoleManagementDirectoryRoleAssignmentAppScopeClient{
		Client: client,
	}, nil
}
