package entitlementmanagementresourcerequestcatalogresourcescope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementResourceRequestCatalogResourceScopeClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementResourceRequestCatalogResourceScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementResourceRequestCatalogResourceScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementresourcerequestcatalogresourcescope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementResourceRequestCatalogResourceScopeClient: %+v", err)
	}

	return &EntitlementManagementResourceRequestCatalogResourceScopeClient{
		Client: client,
	}, nil
}
