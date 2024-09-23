package entitlementmanagementaccesspackagecatalogaccesspackageresourceroleaccesspackageresourceaccesspackageresourceenvironment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageCatalogAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceEnvironmentClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageCatalogAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceEnvironmentClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageCatalogAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceEnvironmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackagecatalogaccesspackageresourceroleaccesspackageresourceaccesspackageresourceenvironment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageCatalogAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceEnvironmentClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageCatalogAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceEnvironmentClient{
		Client: client,
	}, nil
}
