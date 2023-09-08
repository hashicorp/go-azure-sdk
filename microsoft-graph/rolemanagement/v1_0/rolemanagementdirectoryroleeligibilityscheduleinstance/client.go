package rolemanagementdirectoryroleeligibilityscheduleinstance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementDirectoryRoleEligibilityScheduleInstanceClient struct {
	Client *msgraph.Client
}

func NewRoleManagementDirectoryRoleEligibilityScheduleInstanceClientWithBaseURI(api sdkEnv.Api) (*RoleManagementDirectoryRoleEligibilityScheduleInstanceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagementdirectoryroleeligibilityscheduleinstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementDirectoryRoleEligibilityScheduleInstanceClient: %+v", err)
	}

	return &RoleManagementDirectoryRoleEligibilityScheduleInstanceClient{
		Client: client,
	}, nil
}
