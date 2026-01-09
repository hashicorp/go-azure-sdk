package entitlementmanagementroleassignmentscheduleprincipal

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementRoleAssignmentSchedulePrincipalClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementRoleAssignmentSchedulePrincipalClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementRoleAssignmentSchedulePrincipalClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementroleassignmentscheduleprincipal", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementRoleAssignmentSchedulePrincipalClient: %+v", err)
	}

	return &EntitlementManagementRoleAssignmentSchedulePrincipalClient{
		Client: client,
	}, nil
}
