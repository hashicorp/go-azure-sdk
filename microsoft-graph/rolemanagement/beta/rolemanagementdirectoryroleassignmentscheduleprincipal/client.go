package rolemanagementdirectoryroleassignmentscheduleprincipal

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementDirectoryRoleAssignmentSchedulePrincipalClient struct {
	Client *msgraph.Client
}

func NewRoleManagementDirectoryRoleAssignmentSchedulePrincipalClientWithBaseURI(api sdkEnv.Api) (*RoleManagementDirectoryRoleAssignmentSchedulePrincipalClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagementdirectoryroleassignmentscheduleprincipal", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementDirectoryRoleAssignmentSchedulePrincipalClient: %+v", err)
	}

	return &RoleManagementDirectoryRoleAssignmentSchedulePrincipalClient{
		Client: client,
	}, nil
}
