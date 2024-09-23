package entitlementmanagementaccesspackagecatalogaccesspackageresourceroleaccesspackageresource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageCatalogAccessPackageResourceRoleAccessPackageResourceClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageCatalogAccessPackageResourceRoleAccessPackageResourceClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageCatalogAccessPackageResourceRoleAccessPackageResourceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackagecatalogaccesspackageresourceroleaccesspackageresource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageCatalogAccessPackageResourceRoleAccessPackageResourceClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageCatalogAccessPackageResourceRoleAccessPackageResourceClient{
		Client: client,
	}, nil
}
