package rolemanagementcloudpcresourcenamespaceresourceaction

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementCloudPCResourceNamespaceResourceActionClient struct {
	Client *msgraph.Client
}

func NewRoleManagementCloudPCResourceNamespaceResourceActionClientWithBaseURI(api sdkEnv.Api) (*RoleManagementCloudPCResourceNamespaceResourceActionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagementcloudpcresourcenamespaceresourceaction", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementCloudPCResourceNamespaceResourceActionClient: %+v", err)
	}

	return &RoleManagementCloudPCResourceNamespaceResourceActionClient{
		Client: client,
	}, nil
}
