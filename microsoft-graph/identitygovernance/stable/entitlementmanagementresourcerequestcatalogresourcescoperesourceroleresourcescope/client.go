package entitlementmanagementresourcerequestcatalogresourcescoperesourceroleresourcescope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementResourceRequestCatalogResourceScopeResourceRoleResourceScopeClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementResourceRequestCatalogResourceScopeResourceRoleResourceScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementResourceRequestCatalogResourceScopeResourceRoleResourceScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementresourcerequestcatalogresourcescoperesourceroleresourcescope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementResourceRequestCatalogResourceScopeResourceRoleResourceScopeClient: %+v", err)
	}

	return &EntitlementManagementResourceRequestCatalogResourceScopeResourceRoleResourceScopeClient{
		Client: client,
	}, nil
}
