package entitlementmanagementresourcerolescopescoperesourceroleresourcescope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementResourceRoleScopeScopeResourceRoleResourceScopeClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementResourceRoleScopeScopeResourceRoleResourceScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementResourceRoleScopeScopeResourceRoleResourceScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementresourcerolescopescoperesourceroleresourcescope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementResourceRoleScopeScopeResourceRoleResourceScopeClient: %+v", err)
	}

	return &EntitlementManagementResourceRoleScopeScopeResourceRoleResourceScopeClient{
		Client: client,
	}, nil
}
