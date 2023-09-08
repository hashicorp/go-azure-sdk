package rolemanagemententitlementmanagementroleassignmentscheduleactivatedusing

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementEntitlementManagementRoleAssignmentScheduleActivatedUsingClient struct {
	Client *msgraph.Client
}

func NewRoleManagementEntitlementManagementRoleAssignmentScheduleActivatedUsingClientWithBaseURI(api sdkEnv.Api) (*RoleManagementEntitlementManagementRoleAssignmentScheduleActivatedUsingClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagemententitlementmanagementroleassignmentscheduleactivatedusing", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementEntitlementManagementRoleAssignmentScheduleActivatedUsingClient: %+v", err)
	}

	return &RoleManagementEntitlementManagementRoleAssignmentScheduleActivatedUsingClient{
		Client: client,
	}, nil
}
