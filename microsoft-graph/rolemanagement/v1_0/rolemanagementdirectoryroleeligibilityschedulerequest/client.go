package rolemanagementdirectoryroleeligibilityschedulerequest

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementDirectoryRoleEligibilityScheduleRequestClient struct {
	Client *msgraph.Client
}

func NewRoleManagementDirectoryRoleEligibilityScheduleRequestClientWithBaseURI(api sdkEnv.Api) (*RoleManagementDirectoryRoleEligibilityScheduleRequestClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagementdirectoryroleeligibilityschedulerequest", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementDirectoryRoleEligibilityScheduleRequestClient: %+v", err)
	}

	return &RoleManagementDirectoryRoleEligibilityScheduleRequestClient{
		Client: client,
	}, nil
}
