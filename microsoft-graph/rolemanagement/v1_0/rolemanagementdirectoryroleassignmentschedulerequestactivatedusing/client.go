package rolemanagementdirectoryroleassignmentschedulerequestactivatedusing

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementDirectoryRoleAssignmentScheduleRequestActivatedUsingClient struct {
	Client *msgraph.Client
}

func NewRoleManagementDirectoryRoleAssignmentScheduleRequestActivatedUsingClientWithBaseURI(api sdkEnv.Api) (*RoleManagementDirectoryRoleAssignmentScheduleRequestActivatedUsingClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagementdirectoryroleassignmentschedulerequestactivatedusing", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementDirectoryRoleAssignmentScheduleRequestActivatedUsingClient: %+v", err)
	}

	return &RoleManagementDirectoryRoleAssignmentScheduleRequestActivatedUsingClient{
		Client: client,
	}, nil
}
