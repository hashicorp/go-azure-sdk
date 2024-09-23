package entitlementmanagementaccesspackageaccesspackagecatalog

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageAccessPackageCatalogClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageAccessPackageCatalogClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageAccessPackageCatalogClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackageaccesspackagecatalog", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageAccessPackageCatalogClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageAccessPackageCatalogClient{
		Client: client,
	}, nil
}
