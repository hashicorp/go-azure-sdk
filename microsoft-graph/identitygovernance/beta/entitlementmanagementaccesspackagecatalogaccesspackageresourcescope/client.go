package entitlementmanagementaccesspackagecatalogaccesspackageresourcescope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageCatalogAccessPackageResourceScopeClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageCatalogAccessPackageResourceScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageCatalogAccessPackageResourceScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackagecatalogaccesspackageresourcescope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageCatalogAccessPackageResourceScopeClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageCatalogAccessPackageResourceScopeClient{
		Client: client,
	}, nil
}
