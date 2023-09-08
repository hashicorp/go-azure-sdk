package rolemanagementdirectoryroleeligibilityscheduleinstanceappscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementDirectoryRoleEligibilityScheduleInstanceAppScopeClient struct {
	Client *msgraph.Client
}

func NewRoleManagementDirectoryRoleEligibilityScheduleInstanceAppScopeClientWithBaseURI(api sdkEnv.Api) (*RoleManagementDirectoryRoleEligibilityScheduleInstanceAppScopeClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagementdirectoryroleeligibilityscheduleinstanceappscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementDirectoryRoleEligibilityScheduleInstanceAppScopeClient: %+v", err)
	}

	return &RoleManagementDirectoryRoleEligibilityScheduleInstanceAppScopeClient{
		Client: client,
	}, nil
}
