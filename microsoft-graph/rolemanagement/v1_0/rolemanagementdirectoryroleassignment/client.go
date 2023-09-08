package rolemanagementdirectoryroleassignment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementDirectoryRoleAssignmentClient struct {
	Client *msgraph.Client
}

func NewRoleManagementDirectoryRoleAssignmentClientWithBaseURI(api sdkEnv.Api) (*RoleManagementDirectoryRoleAssignmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagementdirectoryroleassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementDirectoryRoleAssignmentClient: %+v", err)
	}

	return &RoleManagementDirectoryRoleAssignmentClient{
		Client: client,
	}, nil
}
