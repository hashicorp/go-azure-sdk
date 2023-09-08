package rolemanagementdirectoryresourcenamespaceresourceactionauthenticationcontext

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementDirectoryResourceNamespaceResourceActionAuthenticationContextClient struct {
	Client *msgraph.Client
}

func NewRoleManagementDirectoryResourceNamespaceResourceActionAuthenticationContextClientWithBaseURI(api sdkEnv.Api) (*RoleManagementDirectoryResourceNamespaceResourceActionAuthenticationContextClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagementdirectoryresourcenamespaceresourceactionauthenticationcontext", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementDirectoryResourceNamespaceResourceActionAuthenticationContextClient: %+v", err)
	}

	return &RoleManagementDirectoryResourceNamespaceResourceActionAuthenticationContextClient{
		Client: client,
	}, nil
}
