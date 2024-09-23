package entitlementmanagementresourcerequestcatalogresourceroleresource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementResourceRequestCatalogResourceRoleResourceClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementResourceRequestCatalogResourceRoleResourceClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementResourceRequestCatalogResourceRoleResourceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementresourcerequestcatalogresourceroleresource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementResourceRequestCatalogResourceRoleResourceClient: %+v", err)
	}

	return &EntitlementManagementResourceRequestCatalogResourceRoleResourceClient{
		Client: client,
	}, nil
}
