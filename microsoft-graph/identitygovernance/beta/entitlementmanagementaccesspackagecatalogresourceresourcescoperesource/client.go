package entitlementmanagementaccesspackagecatalogresourceresourcescoperesource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageCatalogResourceResourceScopeResourceClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageCatalogResourceResourceScopeResourceClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageCatalogResourceResourceScopeResourceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackagecatalogresourceresourcescoperesource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageCatalogResourceResourceScopeResourceClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageCatalogResourceResourceScopeResourceClient{
		Client: client,
	}, nil
}
