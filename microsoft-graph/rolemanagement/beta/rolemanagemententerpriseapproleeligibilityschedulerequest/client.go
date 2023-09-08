package rolemanagemententerpriseapproleeligibilityschedulerequest

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementEnterpriseAppRoleEligibilityScheduleRequestClient struct {
	Client *msgraph.Client
}

func NewRoleManagementEnterpriseAppRoleEligibilityScheduleRequestClientWithBaseURI(api sdkEnv.Api) (*RoleManagementEnterpriseAppRoleEligibilityScheduleRequestClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagemententerpriseapproleeligibilityschedulerequest", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementEnterpriseAppRoleEligibilityScheduleRequestClient: %+v", err)
	}

	return &RoleManagementEnterpriseAppRoleEligibilityScheduleRequestClient{
		Client: client,
	}, nil
}
