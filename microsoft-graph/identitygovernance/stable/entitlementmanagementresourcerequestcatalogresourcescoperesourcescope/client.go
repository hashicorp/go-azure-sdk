package entitlementmanagementresourcerequestcatalogresourcescoperesourcescope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementResourceRequestCatalogResourceScopeResourceScopeClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementResourceRequestCatalogResourceScopeResourceScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementResourceRequestCatalogResourceScopeResourceScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementresourcerequestcatalogresourcescoperesourcescope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementResourceRequestCatalogResourceScopeResourceScopeClient: %+v", err)
	}

	return &EntitlementManagementResourceRequestCatalogResourceScopeResourceScopeClient{
		Client: client,
	}, nil
}
