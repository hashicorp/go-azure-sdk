package entitlementmanagementroledefinitioninheritspermissionsfrom

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementRoleDefinitionInheritsPermissionsFromClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementRoleDefinitionInheritsPermissionsFromClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementRoleDefinitionInheritsPermissionsFromClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementroledefinitioninheritspermissionsfrom", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementRoleDefinitionInheritsPermissionsFromClient: %+v", err)
	}

	return &EntitlementManagementRoleDefinitionInheritsPermissionsFromClient{
		Client: client,
	}, nil
}
