package entitlementmanagementroleassignmentscheduleactivatedusing

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementRoleAssignmentScheduleActivatedUsingClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementRoleAssignmentScheduleActivatedUsingClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementRoleAssignmentScheduleActivatedUsingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementroleassignmentscheduleactivatedusing", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementRoleAssignmentScheduleActivatedUsingClient: %+v", err)
	}

	return &EntitlementManagementRoleAssignmentScheduleActivatedUsingClient{
		Client: client,
	}, nil
}
