package rolemanagementexchangeresourcenamespaceresourceactionauthenticationcontext

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementExchangeResourceNamespaceResourceActionAuthenticationContextClient struct {
	Client *msgraph.Client
}

func NewRoleManagementExchangeResourceNamespaceResourceActionAuthenticationContextClientWithBaseURI(api sdkEnv.Api) (*RoleManagementExchangeResourceNamespaceResourceActionAuthenticationContextClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagementexchangeresourcenamespaceresourceactionauthenticationcontext", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementExchangeResourceNamespaceResourceActionAuthenticationContextClient: %+v", err)
	}

	return &RoleManagementExchangeResourceNamespaceResourceActionAuthenticationContextClient{
		Client: client,
	}, nil
}
