package rolemanagemententitlementmanagementroleeligibilityscheduleprincipal

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementEntitlementManagementRoleEligibilitySchedulePrincipalClient struct {
	Client *msgraph.Client
}

func NewRoleManagementEntitlementManagementRoleEligibilitySchedulePrincipalClientWithBaseURI(api sdkEnv.Api) (*RoleManagementEntitlementManagementRoleEligibilitySchedulePrincipalClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagemententitlementmanagementroleeligibilityscheduleprincipal", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementEntitlementManagementRoleEligibilitySchedulePrincipalClient: %+v", err)
	}

	return &RoleManagementEntitlementManagementRoleEligibilitySchedulePrincipalClient{
		Client: client,
	}, nil
}
