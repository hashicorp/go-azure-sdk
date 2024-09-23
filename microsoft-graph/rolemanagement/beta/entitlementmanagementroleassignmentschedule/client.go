package entitlementmanagementroleassignmentschedule

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementRoleAssignmentScheduleClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementRoleAssignmentScheduleClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementRoleAssignmentScheduleClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementroleassignmentschedule", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementRoleAssignmentScheduleClient: %+v", err)
	}

	return &EntitlementManagementRoleAssignmentScheduleClient{
		Client: client,
	}, nil
}
