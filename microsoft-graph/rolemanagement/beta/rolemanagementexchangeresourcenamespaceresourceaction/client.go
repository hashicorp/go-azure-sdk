package rolemanagementexchangeresourcenamespaceresourceaction

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementExchangeResourceNamespaceResourceActionClient struct {
	Client *msgraph.Client
}

func NewRoleManagementExchangeResourceNamespaceResourceActionClientWithBaseURI(api sdkEnv.Api) (*RoleManagementExchangeResourceNamespaceResourceActionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagementexchangeresourcenamespaceresourceaction", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementExchangeResourceNamespaceResourceActionClient: %+v", err)
	}

	return &RoleManagementExchangeResourceNamespaceResourceActionClient{
		Client: client,
	}, nil
}
