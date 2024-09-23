package entitlementmanagementresourcerequestcatalogresourceroleresourcescoperesourcerole

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceRoleClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceRoleClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceRoleClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementresourcerequestcatalogresourceroleresourcescoperesourcerole", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceRoleClient: %+v", err)
	}

	return &EntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceRoleClient{
		Client: client,
	}, nil
}
