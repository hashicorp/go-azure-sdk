package entitlementmanagementresourcerequestcatalogresourceroleresourcescope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementResourceRequestCatalogResourceRoleResourceScopeClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementResourceRequestCatalogResourceRoleResourceScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementResourceRequestCatalogResourceRoleResourceScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementresourcerequestcatalogresourceroleresourcescope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementResourceRequestCatalogResourceRoleResourceScopeClient: %+v", err)
	}

	return &EntitlementManagementResourceRequestCatalogResourceRoleResourceScopeClient{
		Client: client,
	}, nil
}
