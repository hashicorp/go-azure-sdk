package entitlementmanagementcatalogaccesspackage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementCatalogAccessPackageClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementCatalogAccessPackageClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementCatalogAccessPackageClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementcatalogaccesspackage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementCatalogAccessPackageClient: %+v", err)
	}

	return &EntitlementManagementCatalogAccessPackageClient{
		Client: client,
	}, nil
}
