package entitlementmanagementroleassignmentscheduleinstanceprincipal

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementRoleAssignmentScheduleInstancePrincipalClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementRoleAssignmentScheduleInstancePrincipalClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementRoleAssignmentScheduleInstancePrincipalClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementroleassignmentscheduleinstanceprincipal", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementRoleAssignmentScheduleInstancePrincipalClient: %+v", err)
	}

	return &EntitlementManagementRoleAssignmentScheduleInstancePrincipalClient{
		Client: client,
	}, nil
}
