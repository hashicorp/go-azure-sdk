package rolemanagementdirectoryroleassignmentschedulerequestprincipal

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementDirectoryRoleAssignmentScheduleRequestPrincipalClient struct {
	Client *msgraph.Client
}

func NewRoleManagementDirectoryRoleAssignmentScheduleRequestPrincipalClientWithBaseURI(api sdkEnv.Api) (*RoleManagementDirectoryRoleAssignmentScheduleRequestPrincipalClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagementdirectoryroleassignmentschedulerequestprincipal", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementDirectoryRoleAssignmentScheduleRequestPrincipalClient: %+v", err)
	}

	return &RoleManagementDirectoryRoleAssignmentScheduleRequestPrincipalClient{
		Client: client,
	}, nil
}
