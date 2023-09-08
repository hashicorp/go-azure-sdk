package rolemanagemententitlementmanagementresourcenamespaceresourceactionauthenticationcontext

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementEntitlementManagementResourceNamespaceResourceActionAuthenticationContextClient struct {
	Client *msgraph.Client
}

func NewRoleManagementEntitlementManagementResourceNamespaceResourceActionAuthenticationContextClientWithBaseURI(api sdkEnv.Api) (*RoleManagementEntitlementManagementResourceNamespaceResourceActionAuthenticationContextClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagemententitlementmanagementresourcenamespaceresourceactionauthenticationcontext", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementEntitlementManagementResourceNamespaceResourceActionAuthenticationContextClient: %+v", err)
	}

	return &RoleManagementEntitlementManagementResourceNamespaceResourceActionAuthenticationContextClient{
		Client: client,
	}, nil
}
