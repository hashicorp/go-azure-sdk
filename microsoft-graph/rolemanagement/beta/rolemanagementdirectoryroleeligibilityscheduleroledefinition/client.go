package rolemanagementdirectoryroleeligibilityscheduleroledefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementDirectoryRoleEligibilityScheduleRoleDefinitionClient struct {
	Client *msgraph.Client
}

func NewRoleManagementDirectoryRoleEligibilityScheduleRoleDefinitionClientWithBaseURI(api sdkEnv.Api) (*RoleManagementDirectoryRoleEligibilityScheduleRoleDefinitionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagementdirectoryroleeligibilityscheduleroledefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementDirectoryRoleEligibilityScheduleRoleDefinitionClient: %+v", err)
	}

	return &RoleManagementDirectoryRoleEligibilityScheduleRoleDefinitionClient{
		Client: client,
	}, nil
}
