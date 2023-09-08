package rolemanagementdirectorytransitiveroleassignmentprincipal

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementDirectoryTransitiveRoleAssignmentPrincipalClient struct {
	Client *msgraph.Client
}

func NewRoleManagementDirectoryTransitiveRoleAssignmentPrincipalClientWithBaseURI(api sdkEnv.Api) (*RoleManagementDirectoryTransitiveRoleAssignmentPrincipalClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagementdirectorytransitiveroleassignmentprincipal", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementDirectoryTransitiveRoleAssignmentPrincipalClient: %+v", err)
	}

	return &RoleManagementDirectoryTransitiveRoleAssignmentPrincipalClient{
		Client: client,
	}, nil
}
