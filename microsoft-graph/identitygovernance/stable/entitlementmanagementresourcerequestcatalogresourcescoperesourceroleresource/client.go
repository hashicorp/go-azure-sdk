package entitlementmanagementresourcerequestcatalogresourcescoperesourceroleresource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementResourceRequestCatalogResourceScopeResourceRoleResourceClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementResourceRequestCatalogResourceScopeResourceRoleResourceClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementResourceRequestCatalogResourceScopeResourceRoleResourceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementresourcerequestcatalogresourcescoperesourceroleresource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementResourceRequestCatalogResourceScopeResourceRoleResourceClient: %+v", err)
	}

	return &EntitlementManagementResourceRequestCatalogResourceScopeResourceRoleResourceClient{
		Client: client,
	}, nil
}
