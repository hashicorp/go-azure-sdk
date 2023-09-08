package rolemanagementdirectoryresourcenamespaceresourceaction

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementDirectoryResourceNamespaceResourceActionClient struct {
	Client *msgraph.Client
}

func NewRoleManagementDirectoryResourceNamespaceResourceActionClientWithBaseURI(api sdkEnv.Api) (*RoleManagementDirectoryResourceNamespaceResourceActionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagementdirectoryresourcenamespaceresourceaction", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementDirectoryResourceNamespaceResourceActionClient: %+v", err)
	}

	return &RoleManagementDirectoryResourceNamespaceResourceActionClient{
		Client: client,
	}, nil
}
