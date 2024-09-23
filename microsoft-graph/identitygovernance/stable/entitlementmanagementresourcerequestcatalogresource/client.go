package entitlementmanagementresourcerequestcatalogresource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementResourceRequestCatalogResourceClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementResourceRequestCatalogResourceClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementResourceRequestCatalogResourceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementresourcerequestcatalogresource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementResourceRequestCatalogResourceClient: %+v", err)
	}

	return &EntitlementManagementResourceRequestCatalogResourceClient{
		Client: client,
	}, nil
}
