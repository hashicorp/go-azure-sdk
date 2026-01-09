package entitlementmanagementroleassignmentscheduleinstance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementRoleAssignmentScheduleInstanceClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementRoleAssignmentScheduleInstanceClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementRoleAssignmentScheduleInstanceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementroleassignmentscheduleinstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementRoleAssignmentScheduleInstanceClient: %+v", err)
	}

	return &EntitlementManagementRoleAssignmentScheduleInstanceClient{
		Client: client,
	}, nil
}
