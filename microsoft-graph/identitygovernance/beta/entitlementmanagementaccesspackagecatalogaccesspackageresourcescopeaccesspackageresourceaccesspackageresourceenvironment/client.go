package entitlementmanagementaccesspackagecatalogaccesspackageresourcescopeaccesspackageresourceaccesspackageresourceenvironment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageCatalogAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceEnvironmentClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageCatalogAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceEnvironmentClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageCatalogAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceEnvironmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackagecatalogaccesspackageresourcescopeaccesspackageresourceaccesspackageresourceenvironment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageCatalogAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceEnvironmentClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageCatalogAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceEnvironmentClient{
		Client: client,
	}, nil
}
