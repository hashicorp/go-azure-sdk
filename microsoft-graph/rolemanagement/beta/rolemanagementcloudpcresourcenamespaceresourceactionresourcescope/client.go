package rolemanagementcloudpcresourcenamespaceresourceactionresourcescope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementCloudPCResourceNamespaceResourceActionResourceScopeClient struct {
	Client *msgraph.Client
}

func NewRoleManagementCloudPCResourceNamespaceResourceActionResourceScopeClientWithBaseURI(api sdkEnv.Api) (*RoleManagementCloudPCResourceNamespaceResourceActionResourceScopeClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagementcloudpcresourcenamespaceresourceactionresourcescope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementCloudPCResourceNamespaceResourceActionResourceScopeClient: %+v", err)
	}

	return &RoleManagementCloudPCResourceNamespaceResourceActionResourceScopeClient{
		Client: client,
	}, nil
}
