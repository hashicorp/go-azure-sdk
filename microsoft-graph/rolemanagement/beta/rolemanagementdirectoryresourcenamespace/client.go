package rolemanagementdirectoryresourcenamespace

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementDirectoryResourceNamespaceClient struct {
	Client *msgraph.Client
}

func NewRoleManagementDirectoryResourceNamespaceClientWithBaseURI(api sdkEnv.Api) (*RoleManagementDirectoryResourceNamespaceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagementdirectoryresourcenamespace", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementDirectoryResourceNamespaceClient: %+v", err)
	}

	return &RoleManagementDirectoryResourceNamespaceClient{
		Client: client,
	}, nil
}
