package rolemanagemententerpriseapproleeligibilityscheduleinstanceprincipal

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementEnterpriseAppRoleEligibilityScheduleInstancePrincipalClient struct {
	Client *msgraph.Client
}

func NewRoleManagementEnterpriseAppRoleEligibilityScheduleInstancePrincipalClientWithBaseURI(api sdkEnv.Api) (*RoleManagementEnterpriseAppRoleEligibilityScheduleInstancePrincipalClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagemententerpriseapproleeligibilityscheduleinstanceprincipal", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementEnterpriseAppRoleEligibilityScheduleInstancePrincipalClient: %+v", err)
	}

	return &RoleManagementEnterpriseAppRoleEligibilityScheduleInstancePrincipalClient{
		Client: client,
	}, nil
}
