package rolemanagementdirectoryroleassignmentschedulerequesttargetschedule

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementDirectoryRoleAssignmentScheduleRequestTargetScheduleClient struct {
	Client *msgraph.Client
}

func NewRoleManagementDirectoryRoleAssignmentScheduleRequestTargetScheduleClientWithBaseURI(api sdkEnv.Api) (*RoleManagementDirectoryRoleAssignmentScheduleRequestTargetScheduleClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagementdirectoryroleassignmentschedulerequesttargetschedule", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementDirectoryRoleAssignmentScheduleRequestTargetScheduleClient: %+v", err)
	}

	return &RoleManagementDirectoryRoleAssignmentScheduleRequestTargetScheduleClient{
		Client: client,
	}, nil
}
