package entitlementmanagementresourcerolescoperoleresourcescoperesourcerole

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementResourceRoleScopeRoleResourceScopeResourceRoleClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementResourceRoleScopeRoleResourceScopeResourceRoleClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementResourceRoleScopeRoleResourceScopeResourceRoleClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementresourcerolescoperoleresourcescoperesourcerole", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementResourceRoleScopeRoleResourceScopeResourceRoleClient: %+v", err)
	}

	return &EntitlementManagementResourceRoleScopeRoleResourceScopeResourceRoleClient{
		Client: client,
	}, nil
}
