package entitlementmanagementresourceenvironmentresourcescoperesourceroleresource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementResourceEnvironmentResourceScopeResourceRoleResourceClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementResourceEnvironmentResourceScopeResourceRoleResourceClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementResourceEnvironmentResourceScopeResourceRoleResourceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementresourceenvironmentresourcescoperesourceroleresource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementResourceEnvironmentResourceScopeResourceRoleResourceClient: %+v", err)
	}

	return &EntitlementManagementResourceEnvironmentResourceScopeResourceRoleResourceClient{
		Client: client,
	}, nil
}
