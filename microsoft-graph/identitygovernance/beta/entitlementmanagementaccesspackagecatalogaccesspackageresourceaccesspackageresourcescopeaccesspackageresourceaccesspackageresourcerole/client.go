package entitlementmanagementaccesspackagecatalogaccesspackageresourceaccesspackageresourcescopeaccesspackageresourceaccesspackageresourcerole

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageCatalogAccessPackageResourceAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageCatalogAccessPackageResourceAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageCatalogAccessPackageResourceAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackagecatalogaccesspackageresourceaccesspackageresourcescopeaccesspackageresourceaccesspackageresourcerole", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageCatalogAccessPackageResourceAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageCatalogAccessPackageResourceAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleClient{
		Client: client,
	}, nil
}
