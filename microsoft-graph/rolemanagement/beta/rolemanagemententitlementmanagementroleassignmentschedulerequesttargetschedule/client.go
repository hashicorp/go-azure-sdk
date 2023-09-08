package rolemanagemententitlementmanagementroleassignmentschedulerequesttargetschedule

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementEntitlementManagementRoleAssignmentScheduleRequestTargetScheduleClient struct {
	Client *msgraph.Client
}

func NewRoleManagementEntitlementManagementRoleAssignmentScheduleRequestTargetScheduleClientWithBaseURI(api sdkEnv.Api) (*RoleManagementEntitlementManagementRoleAssignmentScheduleRequestTargetScheduleClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagemententitlementmanagementroleassignmentschedulerequesttargetschedule", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementEntitlementManagementRoleAssignmentScheduleRequestTargetScheduleClient: %+v", err)
	}

	return &RoleManagementEntitlementManagementRoleAssignmentScheduleRequestTargetScheduleClient{
		Client: client,
	}, nil
}
