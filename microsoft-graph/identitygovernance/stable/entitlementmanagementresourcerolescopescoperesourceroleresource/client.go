package entitlementmanagementresourcerolescopescoperesourceroleresource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementResourceRoleScopeScopeResourceRoleResourceClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementResourceRoleScopeScopeResourceRoleResourceClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementResourceRoleScopeScopeResourceRoleResourceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementresourcerolescopescoperesourceroleresource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementResourceRoleScopeScopeResourceRoleResourceClient: %+v", err)
	}

	return &EntitlementManagementResourceRoleScopeScopeResourceRoleResourceClient{
		Client: client,
	}, nil
}
