package rolemanagemententerpriseapproleeligibilityschedule

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementEnterpriseAppRoleEligibilityScheduleClient struct {
	Client *msgraph.Client
}

func NewRoleManagementEnterpriseAppRoleEligibilityScheduleClientWithBaseURI(api sdkEnv.Api) (*RoleManagementEnterpriseAppRoleEligibilityScheduleClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagemententerpriseapproleeligibilityschedule", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementEnterpriseAppRoleEligibilityScheduleClient: %+v", err)
	}

	return &RoleManagementEnterpriseAppRoleEligibilityScheduleClient{
		Client: client,
	}, nil
}
