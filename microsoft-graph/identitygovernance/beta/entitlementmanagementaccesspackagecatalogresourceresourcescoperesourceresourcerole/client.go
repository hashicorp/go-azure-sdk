package entitlementmanagementaccesspackagecatalogresourceresourcescoperesourceresourcerole

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageCatalogResourceResourceScopeResourceResourceRoleClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageCatalogResourceResourceScopeResourceResourceRoleClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageCatalogResourceResourceScopeResourceResourceRoleClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackagecatalogresourceresourcescoperesourceresourcerole", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageCatalogResourceResourceScopeResourceResourceRoleClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageCatalogResourceResourceScopeResourceResourceRoleClient{
		Client: client,
	}, nil
}
