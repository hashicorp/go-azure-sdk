package entitlementmanagementcatalogresourcescoperesourceroleresourcescope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementCatalogResourceScopeResourceRoleResourceScopeClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementCatalogResourceScopeResourceRoleResourceScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementCatalogResourceScopeResourceRoleResourceScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementcatalogresourcescoperesourceroleresourcescope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementCatalogResourceScopeResourceRoleResourceScopeClient: %+v", err)
	}

	return &EntitlementManagementCatalogResourceScopeResourceRoleResourceScopeClient{
		Client: client,
	}, nil
}
