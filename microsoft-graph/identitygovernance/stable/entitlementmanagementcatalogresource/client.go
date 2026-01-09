package entitlementmanagementcatalogresource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementCatalogResourceClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementCatalogResourceClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementCatalogResourceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementcatalogresource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementCatalogResourceClient: %+v", err)
	}

	return &EntitlementManagementCatalogResourceClient{
		Client: client,
	}, nil
}
