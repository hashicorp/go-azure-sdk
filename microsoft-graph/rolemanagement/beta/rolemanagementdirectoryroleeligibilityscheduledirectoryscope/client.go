package rolemanagementdirectoryroleeligibilityscheduledirectoryscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementDirectoryRoleEligibilityScheduleDirectoryScopeClient struct {
	Client *msgraph.Client
}

func NewRoleManagementDirectoryRoleEligibilityScheduleDirectoryScopeClientWithBaseURI(api sdkEnv.Api) (*RoleManagementDirectoryRoleEligibilityScheduleDirectoryScopeClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagementdirectoryroleeligibilityscheduledirectoryscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementDirectoryRoleEligibilityScheduleDirectoryScopeClient: %+v", err)
	}

	return &RoleManagementDirectoryRoleEligibilityScheduleDirectoryScopeClient{
		Client: client,
	}, nil
}
