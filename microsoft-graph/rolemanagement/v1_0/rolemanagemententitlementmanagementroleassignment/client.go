package rolemanagemententitlementmanagementroleassignment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementEntitlementManagementRoleAssignmentClient struct {
	Client *msgraph.Client
}

func NewRoleManagementEntitlementManagementRoleAssignmentClientWithBaseURI(api sdkEnv.Api) (*RoleManagementEntitlementManagementRoleAssignmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagemententitlementmanagementroleassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementEntitlementManagementRoleAssignmentClient: %+v", err)
	}

	return &RoleManagementEntitlementManagementRoleAssignmentClient{
		Client: client,
	}, nil
}
