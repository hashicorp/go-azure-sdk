package rolemanagementdirectoryroleassignmentroledefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementDirectoryRoleAssignmentRoleDefinitionClient struct {
	Client *msgraph.Client
}

func NewRoleManagementDirectoryRoleAssignmentRoleDefinitionClientWithBaseURI(api sdkEnv.Api) (*RoleManagementDirectoryRoleAssignmentRoleDefinitionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagementdirectoryroleassignmentroledefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementDirectoryRoleAssignmentRoleDefinitionClient: %+v", err)
	}

	return &RoleManagementDirectoryRoleAssignmentRoleDefinitionClient{
		Client: client,
	}, nil
}
