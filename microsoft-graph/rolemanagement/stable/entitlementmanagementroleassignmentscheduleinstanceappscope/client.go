package entitlementmanagementroleassignmentscheduleinstanceappscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementRoleAssignmentScheduleInstanceAppScopeClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementRoleAssignmentScheduleInstanceAppScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementRoleAssignmentScheduleInstanceAppScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementroleassignmentscheduleinstanceappscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementRoleAssignmentScheduleInstanceAppScopeClient: %+v", err)
	}

	return &EntitlementManagementRoleAssignmentScheduleInstanceAppScopeClient{
		Client: client,
	}, nil
}
