package rolemanagemententitlementmanagementresourcenamespace

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementEntitlementManagementResourceNamespaceClient struct {
	Client *msgraph.Client
}

func NewRoleManagementEntitlementManagementResourceNamespaceClientWithBaseURI(api sdkEnv.Api) (*RoleManagementEntitlementManagementResourceNamespaceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagemententitlementmanagementresourcenamespace", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementEntitlementManagementResourceNamespaceClient: %+v", err)
	}

	return &RoleManagementEntitlementManagementResourceNamespaceClient{
		Client: client,
	}, nil
}
