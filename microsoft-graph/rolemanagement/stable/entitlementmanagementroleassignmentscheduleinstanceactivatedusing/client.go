package entitlementmanagementroleassignmentscheduleinstanceactivatedusing

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementRoleAssignmentScheduleInstanceActivatedUsingClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementRoleAssignmentScheduleInstanceActivatedUsingClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementRoleAssignmentScheduleInstanceActivatedUsingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementroleassignmentscheduleinstanceactivatedusing", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementRoleAssignmentScheduleInstanceActivatedUsingClient: %+v", err)
	}

	return &EntitlementManagementRoleAssignmentScheduleInstanceActivatedUsingClient{
		Client: client,
	}, nil
}
