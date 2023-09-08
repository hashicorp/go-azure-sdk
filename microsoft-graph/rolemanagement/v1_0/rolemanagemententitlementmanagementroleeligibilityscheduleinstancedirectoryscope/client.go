package rolemanagemententitlementmanagementroleeligibilityscheduleinstancedirectoryscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceDirectoryScopeClient struct {
	Client *msgraph.Client
}

func NewRoleManagementEntitlementManagementRoleEligibilityScheduleInstanceDirectoryScopeClientWithBaseURI(api sdkEnv.Api) (*RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceDirectoryScopeClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagemententitlementmanagementroleeligibilityscheduleinstancedirectoryscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceDirectoryScopeClient: %+v", err)
	}

	return &RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceDirectoryScopeClient{
		Client: client,
	}, nil
}
