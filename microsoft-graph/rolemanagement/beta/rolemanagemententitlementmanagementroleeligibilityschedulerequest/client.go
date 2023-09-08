package rolemanagemententitlementmanagementroleeligibilityschedulerequest

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementEntitlementManagementRoleEligibilityScheduleRequestClient struct {
	Client *msgraph.Client
}

func NewRoleManagementEntitlementManagementRoleEligibilityScheduleRequestClientWithBaseURI(api sdkEnv.Api) (*RoleManagementEntitlementManagementRoleEligibilityScheduleRequestClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagemententitlementmanagementroleeligibilityschedulerequest", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementEntitlementManagementRoleEligibilityScheduleRequestClient: %+v", err)
	}

	return &RoleManagementEntitlementManagementRoleEligibilityScheduleRequestClient{
		Client: client,
	}, nil
}
