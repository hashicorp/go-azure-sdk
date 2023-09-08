package rolemanagementdevicemanagementresourcenamespaceresourceactionauthenticationcontext

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementDeviceManagementResourceNamespaceResourceActionAuthenticationContextClient struct {
	Client *msgraph.Client
}

func NewRoleManagementDeviceManagementResourceNamespaceResourceActionAuthenticationContextClientWithBaseURI(api sdkEnv.Api) (*RoleManagementDeviceManagementResourceNamespaceResourceActionAuthenticationContextClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagementdevicemanagementresourcenamespaceresourceactionauthenticationcontext", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementDeviceManagementResourceNamespaceResourceActionAuthenticationContextClient: %+v", err)
	}

	return &RoleManagementDeviceManagementResourceNamespaceResourceActionAuthenticationContextClient{
		Client: client,
	}, nil
}
