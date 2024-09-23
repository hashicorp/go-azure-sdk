package entitlementmanagementroleassignmentschedulerequestappscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementRoleAssignmentScheduleRequestAppScopeClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementRoleAssignmentScheduleRequestAppScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementRoleAssignmentScheduleRequestAppScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementroleassignmentschedulerequestappscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementRoleAssignmentScheduleRequestAppScopeClient: %+v", err)
	}

	return &EntitlementManagementRoleAssignmentScheduleRequestAppScopeClient{
		Client: client,
	}, nil
}
