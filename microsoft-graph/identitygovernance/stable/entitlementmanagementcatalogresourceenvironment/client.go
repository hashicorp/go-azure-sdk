package entitlementmanagementcatalogresourceenvironment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementCatalogResourceEnvironmentClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementCatalogResourceEnvironmentClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementCatalogResourceEnvironmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementcatalogresourceenvironment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementCatalogResourceEnvironmentClient: %+v", err)
	}

	return &EntitlementManagementCatalogResourceEnvironmentClient{
		Client: client,
	}, nil
}
