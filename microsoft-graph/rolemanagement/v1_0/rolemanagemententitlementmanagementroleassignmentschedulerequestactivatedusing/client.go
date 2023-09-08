package rolemanagemententitlementmanagementroleassignmentschedulerequestactivatedusing

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementEntitlementManagementRoleAssignmentScheduleRequestActivatedUsingClient struct {
	Client *msgraph.Client
}

func NewRoleManagementEntitlementManagementRoleAssignmentScheduleRequestActivatedUsingClientWithBaseURI(api sdkEnv.Api) (*RoleManagementEntitlementManagementRoleAssignmentScheduleRequestActivatedUsingClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagemententitlementmanagementroleassignmentschedulerequestactivatedusing", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementEntitlementManagementRoleAssignmentScheduleRequestActivatedUsingClient: %+v", err)
	}

	return &RoleManagementEntitlementManagementRoleAssignmentScheduleRequestActivatedUsingClient{
		Client: client,
	}, nil
}
