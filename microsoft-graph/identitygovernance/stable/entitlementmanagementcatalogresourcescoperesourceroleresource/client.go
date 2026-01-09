package entitlementmanagementcatalogresourcescoperesourceroleresource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementCatalogResourceScopeResourceRoleResourceClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementCatalogResourceScopeResourceRoleResourceClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementCatalogResourceScopeResourceRoleResourceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementcatalogresourcescoperesourceroleresource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementCatalogResourceScopeResourceRoleResourceClient: %+v", err)
	}

	return &EntitlementManagementCatalogResourceScopeResourceRoleResourceClient{
		Client: client,
	}, nil
}
