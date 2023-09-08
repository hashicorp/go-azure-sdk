package rolemanagemententitlementmanagementtransitiveroleassignment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementEntitlementManagementTransitiveRoleAssignmentClient struct {
	Client *msgraph.Client
}

func NewRoleManagementEntitlementManagementTransitiveRoleAssignmentClientWithBaseURI(api sdkEnv.Api) (*RoleManagementEntitlementManagementTransitiveRoleAssignmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagemententitlementmanagementtransitiveroleassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementEntitlementManagementTransitiveRoleAssignmentClient: %+v", err)
	}

	return &RoleManagementEntitlementManagementTransitiveRoleAssignmentClient{
		Client: client,
	}, nil
}
