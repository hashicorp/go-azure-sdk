package entitlementmanagementcatalogresourcescoperesource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementCatalogResourceScopeResourceClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementCatalogResourceScopeResourceClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementCatalogResourceScopeResourceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementcatalogresourcescoperesource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementCatalogResourceScopeResourceClient: %+v", err)
	}

	return &EntitlementManagementCatalogResourceScopeResourceClient{
		Client: client,
	}, nil
}
