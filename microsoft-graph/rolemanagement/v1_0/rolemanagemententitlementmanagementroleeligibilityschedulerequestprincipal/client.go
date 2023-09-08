package rolemanagemententitlementmanagementroleeligibilityschedulerequestprincipal

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementEntitlementManagementRoleEligibilityScheduleRequestPrincipalClient struct {
	Client *msgraph.Client
}

func NewRoleManagementEntitlementManagementRoleEligibilityScheduleRequestPrincipalClientWithBaseURI(api sdkEnv.Api) (*RoleManagementEntitlementManagementRoleEligibilityScheduleRequestPrincipalClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagemententitlementmanagementroleeligibilityschedulerequestprincipal", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementEntitlementManagementRoleEligibilityScheduleRequestPrincipalClient: %+v", err)
	}

	return &RoleManagementEntitlementManagementRoleEligibilityScheduleRequestPrincipalClient{
		Client: client,
	}, nil
}
