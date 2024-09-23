package entitlementmanagementroleassignmentscheduledirectoryscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementRoleAssignmentScheduleDirectoryScopeClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementRoleAssignmentScheduleDirectoryScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementRoleAssignmentScheduleDirectoryScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementroleassignmentscheduledirectoryscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementRoleAssignmentScheduleDirectoryScopeClient: %+v", err)
	}

	return &EntitlementManagementRoleAssignmentScheduleDirectoryScopeClient{
		Client: client,
	}, nil
}
