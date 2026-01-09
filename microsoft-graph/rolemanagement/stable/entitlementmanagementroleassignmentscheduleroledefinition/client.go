package entitlementmanagementroleassignmentscheduleroledefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementRoleAssignmentScheduleRoleDefinitionClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementRoleAssignmentScheduleRoleDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementRoleAssignmentScheduleRoleDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementroleassignmentscheduleroledefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementRoleAssignmentScheduleRoleDefinitionClient: %+v", err)
	}

	return &EntitlementManagementRoleAssignmentScheduleRoleDefinitionClient{
		Client: client,
	}, nil
}
