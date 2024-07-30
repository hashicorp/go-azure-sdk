package entitlementmanagementcatalogresourcescoperesourcerole

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementCatalogResourceScopeResourceRoleClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementCatalogResourceScopeResourceRoleClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementCatalogResourceScopeResourceRoleClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementcatalogresourcescoperesourcerole", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementCatalogResourceScopeResourceRoleClient: %+v", err)
	}

	return &EntitlementManagementCatalogResourceScopeResourceRoleClient{
		Client: client,
	}, nil
}
