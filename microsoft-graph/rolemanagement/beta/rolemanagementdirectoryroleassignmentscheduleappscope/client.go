package rolemanagementdirectoryroleassignmentscheduleappscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementDirectoryRoleAssignmentScheduleAppScopeClient struct {
	Client *msgraph.Client
}

func NewRoleManagementDirectoryRoleAssignmentScheduleAppScopeClientWithBaseURI(api sdkEnv.Api) (*RoleManagementDirectoryRoleAssignmentScheduleAppScopeClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagementdirectoryroleassignmentscheduleappscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementDirectoryRoleAssignmentScheduleAppScopeClient: %+v", err)
	}

	return &RoleManagementDirectoryRoleAssignmentScheduleAppScopeClient{
		Client: client,
	}, nil
}
