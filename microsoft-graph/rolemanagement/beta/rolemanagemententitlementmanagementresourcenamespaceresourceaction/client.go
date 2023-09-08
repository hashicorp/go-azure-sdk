package rolemanagemententitlementmanagementresourcenamespaceresourceaction

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementEntitlementManagementResourceNamespaceResourceActionClient struct {
	Client *msgraph.Client
}

func NewRoleManagementEntitlementManagementResourceNamespaceResourceActionClientWithBaseURI(api sdkEnv.Api) (*RoleManagementEntitlementManagementResourceNamespaceResourceActionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagemententitlementmanagementresourcenamespaceresourceaction", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementEntitlementManagementResourceNamespaceResourceActionClient: %+v", err)
	}

	return &RoleManagementEntitlementManagementResourceNamespaceResourceActionClient{
		Client: client,
	}, nil
}
