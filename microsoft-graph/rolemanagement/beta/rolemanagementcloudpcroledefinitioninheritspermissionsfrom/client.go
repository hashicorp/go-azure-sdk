package rolemanagementcloudpcroledefinitioninheritspermissionsfrom

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementCloudPCRoleDefinitionInheritsPermissionsFromClient struct {
	Client *msgraph.Client
}

func NewRoleManagementCloudPCRoleDefinitionInheritsPermissionsFromClientWithBaseURI(api sdkEnv.Api) (*RoleManagementCloudPCRoleDefinitionInheritsPermissionsFromClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagementcloudpcroledefinitioninheritspermissionsfrom", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementCloudPCRoleDefinitionInheritsPermissionsFromClient: %+v", err)
	}

	return &RoleManagementCloudPCRoleDefinitionInheritsPermissionsFromClient{
		Client: client,
	}, nil
}
