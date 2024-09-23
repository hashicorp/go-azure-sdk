package entitlementmanagementaccesspackagecatalogaccesspackageresourceaccesspackageresourcescopeaccesspackageresourceaccesspackageresourceenvironment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageCatalogAccessPackageResourceAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceEnvironmentClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageCatalogAccessPackageResourceAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceEnvironmentClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageCatalogAccessPackageResourceAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceEnvironmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackagecatalogaccesspackageresourceaccesspackageresourcescopeaccesspackageresourceaccesspackageresourceenvironment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageCatalogAccessPackageResourceAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceEnvironmentClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageCatalogAccessPackageResourceAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceEnvironmentClient{
		Client: client,
	}, nil
}
