package rolemanagemententitlementmanagementroledefinitioninheritspermissionsfrom

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementEntitlementManagementRoleDefinitionInheritsPermissionsFromClient struct {
	Client *msgraph.Client
}

func NewRoleManagementEntitlementManagementRoleDefinitionInheritsPermissionsFromClientWithBaseURI(api sdkEnv.Api) (*RoleManagementEntitlementManagementRoleDefinitionInheritsPermissionsFromClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagemententitlementmanagementroledefinitioninheritspermissionsfrom", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementEntitlementManagementRoleDefinitionInheritsPermissionsFromClient: %+v", err)
	}

	return &RoleManagementEntitlementManagementRoleDefinitionInheritsPermissionsFromClient{
		Client: client,
	}, nil
}
