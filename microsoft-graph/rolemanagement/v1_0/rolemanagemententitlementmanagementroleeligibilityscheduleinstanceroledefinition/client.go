package rolemanagemententitlementmanagementroleeligibilityscheduleinstanceroledefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceRoleDefinitionClient struct {
	Client *msgraph.Client
}

func NewRoleManagementEntitlementManagementRoleEligibilityScheduleInstanceRoleDefinitionClientWithBaseURI(api sdkEnv.Api) (*RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceRoleDefinitionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagemententitlementmanagementroleeligibilityscheduleinstanceroledefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceRoleDefinitionClient: %+v", err)
	}

	return &RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceRoleDefinitionClient{
		Client: client,
	}, nil
}
