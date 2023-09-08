package rolemanagementdirectoryroleassignmentapproval

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementDirectoryRoleAssignmentApprovalClient struct {
	Client *msgraph.Client
}

func NewRoleManagementDirectoryRoleAssignmentApprovalClientWithBaseURI(api sdkEnv.Api) (*RoleManagementDirectoryRoleAssignmentApprovalClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagementdirectoryroleassignmentapproval", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementDirectoryRoleAssignmentApprovalClient: %+v", err)
	}

	return &RoleManagementDirectoryRoleAssignmentApprovalClient{
		Client: client,
	}, nil
}
