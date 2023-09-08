package rolemanagemententitlementmanagementroleassignmentroledefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementEntitlementManagementRoleAssignmentRoleDefinitionClient struct {
	Client *msgraph.Client
}

func NewRoleManagementEntitlementManagementRoleAssignmentRoleDefinitionClientWithBaseURI(api sdkEnv.Api) (*RoleManagementEntitlementManagementRoleAssignmentRoleDefinitionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagemententitlementmanagementroleassignmentroledefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementEntitlementManagementRoleAssignmentRoleDefinitionClient: %+v", err)
	}

	return &RoleManagementEntitlementManagementRoleAssignmentRoleDefinitionClient{
		Client: client,
	}, nil
}
