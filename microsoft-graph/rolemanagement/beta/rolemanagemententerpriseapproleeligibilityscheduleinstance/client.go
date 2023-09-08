package rolemanagemententerpriseapproleeligibilityscheduleinstance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementEnterpriseAppRoleEligibilityScheduleInstanceClient struct {
	Client *msgraph.Client
}

func NewRoleManagementEnterpriseAppRoleEligibilityScheduleInstanceClientWithBaseURI(api sdkEnv.Api) (*RoleManagementEnterpriseAppRoleEligibilityScheduleInstanceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagemententerpriseapproleeligibilityscheduleinstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementEnterpriseAppRoleEligibilityScheduleInstanceClient: %+v", err)
	}

	return &RoleManagementEnterpriseAppRoleEligibilityScheduleInstanceClient{
		Client: client,
	}, nil
}
