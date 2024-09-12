package entitlementmanagementaccesspackagecatalogresourceresourcescoperesourceresourceroleresource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageCatalogResourceResourceScopeResourceResourceRoleResourceClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageCatalogResourceResourceScopeResourceResourceRoleResourceClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageCatalogResourceResourceScopeResourceResourceRoleResourceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackagecatalogresourceresourcescoperesourceresourceroleresource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageCatalogResourceResourceScopeResourceResourceRoleResourceClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageCatalogResourceResourceScopeResourceResourceRoleResourceClient{
		Client: client,
	}, nil
}
