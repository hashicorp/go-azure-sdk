package entitlementmanagementaccesspackagecatalogresourceroleresourceresourcescoperesourceresourceenvironment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageCatalogResourceRoleResourceResourceScopeResourceResourceEnvironmentClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageCatalogResourceRoleResourceResourceScopeResourceResourceEnvironmentClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageCatalogResourceRoleResourceResourceScopeResourceResourceEnvironmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackagecatalogresourceroleresourceresourcescoperesourceresourceenvironment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageCatalogResourceRoleResourceResourceScopeResourceResourceEnvironmentClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageCatalogResourceRoleResourceResourceScopeResourceResourceEnvironmentClient{
		Client: client,
	}, nil
}
