package entitlementmanagementcatalogresourceroleresourceenvironment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementCatalogResourceRoleResourceEnvironmentClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementCatalogResourceRoleResourceEnvironmentClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementCatalogResourceRoleResourceEnvironmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementcatalogresourceroleresourceenvironment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementCatalogResourceRoleResourceEnvironmentClient: %+v", err)
	}

	return &EntitlementManagementCatalogResourceRoleResourceEnvironmentClient{
		Client: client,
	}, nil
}
