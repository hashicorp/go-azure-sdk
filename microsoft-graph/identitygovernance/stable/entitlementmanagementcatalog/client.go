package entitlementmanagementcatalog

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementCatalogClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementCatalogClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementCatalogClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementcatalog", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementCatalogClient: %+v", err)
	}

	return &EntitlementManagementCatalogClient{
		Client: client,
	}, nil
}
