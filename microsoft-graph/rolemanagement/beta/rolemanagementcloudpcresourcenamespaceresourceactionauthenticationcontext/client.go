package rolemanagementcloudpcresourcenamespaceresourceactionauthenticationcontext

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementCloudPCResourceNamespaceResourceActionAuthenticationContextClient struct {
	Client *msgraph.Client
}

func NewRoleManagementCloudPCResourceNamespaceResourceActionAuthenticationContextClientWithBaseURI(api sdkEnv.Api) (*RoleManagementCloudPCResourceNamespaceResourceActionAuthenticationContextClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagementcloudpcresourcenamespaceresourceactionauthenticationcontext", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementCloudPCResourceNamespaceResourceActionAuthenticationContextClient: %+v", err)
	}

	return &RoleManagementCloudPCResourceNamespaceResourceActionAuthenticationContextClient{
		Client: client,
	}, nil
}
