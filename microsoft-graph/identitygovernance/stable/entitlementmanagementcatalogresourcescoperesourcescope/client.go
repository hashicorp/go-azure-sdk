package entitlementmanagementcatalogresourcescoperesourcescope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementCatalogResourceScopeResourceScopeClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementCatalogResourceScopeResourceScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementCatalogResourceScopeResourceScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementcatalogresourcescoperesourcescope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementCatalogResourceScopeResourceScopeClient: %+v", err)
	}

	return &EntitlementManagementCatalogResourceScopeResourceScopeClient{
		Client: client,
	}, nil
}
