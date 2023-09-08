package rolemanagementdirectoryroleassignmentprincipal

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementDirectoryRoleAssignmentPrincipalClient struct {
	Client *msgraph.Client
}

func NewRoleManagementDirectoryRoleAssignmentPrincipalClientWithBaseURI(api sdkEnv.Api) (*RoleManagementDirectoryRoleAssignmentPrincipalClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagementdirectoryroleassignmentprincipal", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementDirectoryRoleAssignmentPrincipalClient: %+v", err)
	}

	return &RoleManagementDirectoryRoleAssignmentPrincipalClient{
		Client: client,
	}, nil
}
