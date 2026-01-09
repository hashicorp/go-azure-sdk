package entitlementmanagementcatalogresourcerole

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementCatalogResourceRoleClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementCatalogResourceRoleClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementCatalogResourceRoleClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementcatalogresourcerole", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementCatalogResourceRoleClient: %+v", err)
	}

	return &EntitlementManagementCatalogResourceRoleClient{
		Client: client,
	}, nil
}
