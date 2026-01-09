package entitlementmanagementtransitiveroleassignmentprincipal

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementTransitiveRoleAssignmentPrincipalClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementTransitiveRoleAssignmentPrincipalClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementTransitiveRoleAssignmentPrincipalClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementtransitiveroleassignmentprincipal", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementTransitiveRoleAssignmentPrincipalClient: %+v", err)
	}

	return &EntitlementManagementTransitiveRoleAssignmentPrincipalClient{
		Client: client,
	}, nil
}
