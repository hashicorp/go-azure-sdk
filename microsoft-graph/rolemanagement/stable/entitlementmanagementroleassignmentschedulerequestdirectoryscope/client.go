package entitlementmanagementroleassignmentschedulerequestdirectoryscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementRoleAssignmentScheduleRequestDirectoryScopeClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementRoleAssignmentScheduleRequestDirectoryScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementRoleAssignmentScheduleRequestDirectoryScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementroleassignmentschedulerequestdirectoryscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementRoleAssignmentScheduleRequestDirectoryScopeClient: %+v", err)
	}

	return &EntitlementManagementRoleAssignmentScheduleRequestDirectoryScopeClient{
		Client: client,
	}, nil
}
