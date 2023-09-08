package rolemanagementdirectoryroleassignmentscheduleinstanceprincipal

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementDirectoryRoleAssignmentScheduleInstancePrincipalClient struct {
	Client *msgraph.Client
}

func NewRoleManagementDirectoryRoleAssignmentScheduleInstancePrincipalClientWithBaseURI(api sdkEnv.Api) (*RoleManagementDirectoryRoleAssignmentScheduleInstancePrincipalClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagementdirectoryroleassignmentscheduleinstanceprincipal", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementDirectoryRoleAssignmentScheduleInstancePrincipalClient: %+v", err)
	}

	return &RoleManagementDirectoryRoleAssignmentScheduleInstancePrincipalClient{
		Client: client,
	}, nil
}
