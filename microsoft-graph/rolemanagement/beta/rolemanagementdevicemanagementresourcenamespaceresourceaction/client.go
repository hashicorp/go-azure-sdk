package rolemanagementdevicemanagementresourcenamespaceresourceaction

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementDeviceManagementResourceNamespaceResourceActionClient struct {
	Client *msgraph.Client
}

func NewRoleManagementDeviceManagementResourceNamespaceResourceActionClientWithBaseURI(api sdkEnv.Api) (*RoleManagementDeviceManagementResourceNamespaceResourceActionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagementdevicemanagementresourcenamespaceresourceaction", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementDeviceManagementResourceNamespaceResourceActionClient: %+v", err)
	}

	return &RoleManagementDeviceManagementResourceNamespaceResourceActionClient{
		Client: client,
	}, nil
}
