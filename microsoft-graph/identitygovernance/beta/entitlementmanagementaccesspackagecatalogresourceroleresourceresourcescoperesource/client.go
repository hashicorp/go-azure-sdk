package entitlementmanagementaccesspackagecatalogresourceroleresourceresourcescoperesource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageCatalogResourceRoleResourceResourceScopeResourceClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageCatalogResourceRoleResourceResourceScopeResourceClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageCatalogResourceRoleResourceResourceScopeResourceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackagecatalogresourceroleresourceresourcescoperesource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageCatalogResourceRoleResourceResourceScopeResourceClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageCatalogResourceRoleResourceResourceScopeResourceClient{
		Client: client,
	}, nil
}
