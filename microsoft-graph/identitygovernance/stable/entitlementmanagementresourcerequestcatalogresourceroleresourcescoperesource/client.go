package entitlementmanagementresourcerequestcatalogresourceroleresourcescoperesource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementresourcerequestcatalogresourceroleresourcescoperesource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceClient: %+v", err)
	}

	return &EntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceClient{
		Client: client,
	}, nil
}
