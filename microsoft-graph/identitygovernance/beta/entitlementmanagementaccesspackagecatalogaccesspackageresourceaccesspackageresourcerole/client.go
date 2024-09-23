package entitlementmanagementaccesspackagecatalogaccesspackageresourceaccesspackageresourcerole

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageCatalogAccessPackageResourceAccessPackageResourceRoleClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageCatalogAccessPackageResourceAccessPackageResourceRoleClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageCatalogAccessPackageResourceAccessPackageResourceRoleClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackagecatalogaccesspackageresourceaccesspackageresourcerole", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageCatalogAccessPackageResourceAccessPackageResourceRoleClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageCatalogAccessPackageResourceAccessPackageResourceRoleClient{
		Client: client,
	}, nil
}
