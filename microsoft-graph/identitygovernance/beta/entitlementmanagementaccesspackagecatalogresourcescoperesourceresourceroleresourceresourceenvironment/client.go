package entitlementmanagementaccesspackagecatalogresourcescoperesourceresourceroleresourceresourceenvironment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageCatalogResourceScopeResourceResourceRoleResourceResourceEnvironmentClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageCatalogResourceScopeResourceResourceRoleResourceResourceEnvironmentClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageCatalogResourceScopeResourceResourceRoleResourceResourceEnvironmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackagecatalogresourcescoperesourceresourceroleresourceresourceenvironment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageCatalogResourceScopeResourceResourceRoleResourceResourceEnvironmentClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageCatalogResourceScopeResourceResourceRoleResourceResourceEnvironmentClient{
		Client: client,
	}, nil
}
