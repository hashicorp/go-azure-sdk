package rolemanagemententerpriseapproledefinitioninheritspermissionsfrom

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementEnterpriseAppRoleDefinitionInheritsPermissionsFromClient struct {
	Client *msgraph.Client
}

func NewRoleManagementEnterpriseAppRoleDefinitionInheritsPermissionsFromClientWithBaseURI(api sdkEnv.Api) (*RoleManagementEnterpriseAppRoleDefinitionInheritsPermissionsFromClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagemententerpriseapproledefinitioninheritspermissionsfrom", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementEnterpriseAppRoleDefinitionInheritsPermissionsFromClient: %+v", err)
	}

	return &RoleManagementEnterpriseAppRoleDefinitionInheritsPermissionsFromClient{
		Client: client,
	}, nil
}
