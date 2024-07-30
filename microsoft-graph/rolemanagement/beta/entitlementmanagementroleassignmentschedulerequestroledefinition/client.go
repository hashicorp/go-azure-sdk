package entitlementmanagementroleassignmentschedulerequestroledefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementRoleAssignmentScheduleRequestRoleDefinitionClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementRoleAssignmentScheduleRequestRoleDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementRoleAssignmentScheduleRequestRoleDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementroleassignmentschedulerequestroledefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementRoleAssignmentScheduleRequestRoleDefinitionClient: %+v", err)
	}

	return &EntitlementManagementRoleAssignmentScheduleRequestRoleDefinitionClient{
		Client: client,
	}, nil
}
