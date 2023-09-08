package rolemanagemententitlementmanagementroleassignmentschedulerequestappscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementEntitlementManagementRoleAssignmentScheduleRequestAppScopeClient struct {
	Client *msgraph.Client
}

func NewRoleManagementEntitlementManagementRoleAssignmentScheduleRequestAppScopeClientWithBaseURI(api sdkEnv.Api) (*RoleManagementEntitlementManagementRoleAssignmentScheduleRequestAppScopeClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagemententitlementmanagementroleassignmentschedulerequestappscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementEntitlementManagementRoleAssignmentScheduleRequestAppScopeClient: %+v", err)
	}

	return &RoleManagementEntitlementManagementRoleAssignmentScheduleRequestAppScopeClient{
		Client: client,
	}, nil
}
