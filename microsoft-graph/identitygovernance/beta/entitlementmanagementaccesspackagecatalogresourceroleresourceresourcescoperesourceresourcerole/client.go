package entitlementmanagementaccesspackagecatalogresourceroleresourceresourcescoperesourceresourcerole

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageCatalogResourceRoleResourceResourceScopeResourceResourceRoleClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageCatalogResourceRoleResourceResourceScopeResourceResourceRoleClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageCatalogResourceRoleResourceResourceScopeResourceResourceRoleClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackagecatalogresourceroleresourceresourcescoperesourceresourcerole", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageCatalogResourceRoleResourceResourceScopeResourceResourceRoleClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageCatalogResourceRoleResourceResourceScopeResourceResourceRoleClient{
		Client: client,
	}, nil
}
