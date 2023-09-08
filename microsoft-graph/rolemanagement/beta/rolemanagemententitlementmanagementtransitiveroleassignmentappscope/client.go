package rolemanagemententitlementmanagementtransitiveroleassignmentappscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementEntitlementManagementTransitiveRoleAssignmentAppScopeClient struct {
	Client *msgraph.Client
}

func NewRoleManagementEntitlementManagementTransitiveRoleAssignmentAppScopeClientWithBaseURI(api sdkEnv.Api) (*RoleManagementEntitlementManagementTransitiveRoleAssignmentAppScopeClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagemententitlementmanagementtransitiveroleassignmentappscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementEntitlementManagementTransitiveRoleAssignmentAppScopeClient: %+v", err)
	}

	return &RoleManagementEntitlementManagementTransitiveRoleAssignmentAppScopeClient{
		Client: client,
	}, nil
}
