package rolemanagemententitlementmanagementroleeligibilityschedule

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementEntitlementManagementRoleEligibilityScheduleClient struct {
	Client *msgraph.Client
}

func NewRoleManagementEntitlementManagementRoleEligibilityScheduleClientWithBaseURI(api sdkEnv.Api) (*RoleManagementEntitlementManagementRoleEligibilityScheduleClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagemententitlementmanagementroleeligibilityschedule", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementEntitlementManagementRoleEligibilityScheduleClient: %+v", err)
	}

	return &RoleManagementEntitlementManagementRoleEligibilityScheduleClient{
		Client: client,
	}, nil
}
