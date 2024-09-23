package entitlementmanagementresourcerolescoperoleresourcerole

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementResourceRoleScopeRoleResourceRoleClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementResourceRoleScopeRoleResourceRoleClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementResourceRoleScopeRoleResourceRoleClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementresourcerolescoperoleresourcerole", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementResourceRoleScopeRoleResourceRoleClient: %+v", err)
	}

	return &EntitlementManagementResourceRoleScopeRoleResourceRoleClient{
		Client: client,
	}, nil
}
