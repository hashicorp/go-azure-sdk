package entitlementmanagementresourceenvironmentresourceroleresourcescoperesource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementResourceEnvironmentResourceRoleResourceScopeResourceClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementResourceEnvironmentResourceRoleResourceScopeResourceClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementResourceEnvironmentResourceRoleResourceScopeResourceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementresourceenvironmentresourceroleresourcescoperesource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementResourceEnvironmentResourceRoleResourceScopeResourceClient: %+v", err)
	}

	return &EntitlementManagementResourceEnvironmentResourceRoleResourceScopeResourceClient{
		Client: client,
	}, nil
}
