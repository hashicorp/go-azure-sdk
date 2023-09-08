package rolemanagementdirectory

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementDirectoryClient struct {
	Client *msgraph.Client
}

func NewRoleManagementDirectoryClientWithBaseURI(api sdkEnv.Api) (*RoleManagementDirectoryClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagementdirectory", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementDirectoryClient: %+v", err)
	}

	return &RoleManagementDirectoryClient{
		Client: client,
	}, nil
}
