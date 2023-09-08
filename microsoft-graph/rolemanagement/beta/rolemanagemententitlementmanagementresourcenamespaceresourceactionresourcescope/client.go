package rolemanagemententitlementmanagementresourcenamespaceresourceactionresourcescope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementEntitlementManagementResourceNamespaceResourceActionResourceScopeClient struct {
	Client *msgraph.Client
}

func NewRoleManagementEntitlementManagementResourceNamespaceResourceActionResourceScopeClientWithBaseURI(api sdkEnv.Api) (*RoleManagementEntitlementManagementResourceNamespaceResourceActionResourceScopeClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagemententitlementmanagementresourcenamespaceresourceactionresourcescope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementEntitlementManagementResourceNamespaceResourceActionResourceScopeClient: %+v", err)
	}

	return &RoleManagementEntitlementManagementResourceNamespaceResourceActionResourceScopeClient{
		Client: client,
	}, nil
}
