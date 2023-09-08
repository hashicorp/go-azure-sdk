package rolemanagementdirectoryroleeligibilityscheduleprincipal

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementDirectoryRoleEligibilitySchedulePrincipalClient struct {
	Client *msgraph.Client
}

func NewRoleManagementDirectoryRoleEligibilitySchedulePrincipalClientWithBaseURI(api sdkEnv.Api) (*RoleManagementDirectoryRoleEligibilitySchedulePrincipalClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagementdirectoryroleeligibilityscheduleprincipal", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementDirectoryRoleEligibilitySchedulePrincipalClient: %+v", err)
	}

	return &RoleManagementDirectoryRoleEligibilitySchedulePrincipalClient{
		Client: client,
	}, nil
}
