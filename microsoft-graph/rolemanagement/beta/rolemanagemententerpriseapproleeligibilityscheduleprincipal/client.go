package rolemanagemententerpriseapproleeligibilityscheduleprincipal

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementEnterpriseAppRoleEligibilitySchedulePrincipalClient struct {
	Client *msgraph.Client
}

func NewRoleManagementEnterpriseAppRoleEligibilitySchedulePrincipalClientWithBaseURI(api sdkEnv.Api) (*RoleManagementEnterpriseAppRoleEligibilitySchedulePrincipalClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagemententerpriseapproleeligibilityscheduleprincipal", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementEnterpriseAppRoleEligibilitySchedulePrincipalClient: %+v", err)
	}

	return &RoleManagementEnterpriseAppRoleEligibilitySchedulePrincipalClient{
		Client: client,
	}, nil
}
