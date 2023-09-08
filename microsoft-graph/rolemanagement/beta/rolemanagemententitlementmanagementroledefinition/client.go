package rolemanagemententitlementmanagementroledefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementEntitlementManagementRoleDefinitionClient struct {
	Client *msgraph.Client
}

func NewRoleManagementEntitlementManagementRoleDefinitionClientWithBaseURI(api sdkEnv.Api) (*RoleManagementEntitlementManagementRoleDefinitionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagemententitlementmanagementroledefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementEntitlementManagementRoleDefinitionClient: %+v", err)
	}

	return &RoleManagementEntitlementManagementRoleDefinitionClient{
		Client: client,
	}, nil
}
