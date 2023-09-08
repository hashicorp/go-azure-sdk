package rolemanagementdirectoryresourcenamespaceresourceactionresourcescope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementDirectoryResourceNamespaceResourceActionResourceScopeClient struct {
	Client *msgraph.Client
}

func NewRoleManagementDirectoryResourceNamespaceResourceActionResourceScopeClientWithBaseURI(api sdkEnv.Api) (*RoleManagementDirectoryResourceNamespaceResourceActionResourceScopeClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagementdirectoryresourcenamespaceresourceactionresourcescope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementDirectoryResourceNamespaceResourceActionResourceScopeClient: %+v", err)
	}

	return &RoleManagementDirectoryResourceNamespaceResourceActionResourceScopeClient{
		Client: client,
	}, nil
}
