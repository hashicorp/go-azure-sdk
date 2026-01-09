package entitlementmanagementresourcerolescoperoleresourcescope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementResourceRoleScopeRoleResourceScopeClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementResourceRoleScopeRoleResourceScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementResourceRoleScopeRoleResourceScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementresourcerolescoperoleresourcescope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementResourceRoleScopeRoleResourceScopeClient: %+v", err)
	}

	return &EntitlementManagementResourceRoleScopeRoleResourceScopeClient{
		Client: client,
	}, nil
}
