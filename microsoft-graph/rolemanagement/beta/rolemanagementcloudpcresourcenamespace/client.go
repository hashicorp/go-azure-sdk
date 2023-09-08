package rolemanagementcloudpcresourcenamespace

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementCloudPCResourceNamespaceClient struct {
	Client *msgraph.Client
}

func NewRoleManagementCloudPCResourceNamespaceClientWithBaseURI(api sdkEnv.Api) (*RoleManagementCloudPCResourceNamespaceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagementcloudpcresourcenamespace", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementCloudPCResourceNamespaceClient: %+v", err)
	}

	return &RoleManagementCloudPCResourceNamespaceClient{
		Client: client,
	}, nil
}
