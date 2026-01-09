package entitlementmanagementaccesspackagecatalogcustomaccesspackageworkflowextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageCatalogCustomAccessPackageWorkflowExtensionClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageCatalogCustomAccessPackageWorkflowExtensionClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageCatalogCustomAccessPackageWorkflowExtensionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackagecatalogcustomaccesspackageworkflowextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageCatalogCustomAccessPackageWorkflowExtensionClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageCatalogCustomAccessPackageWorkflowExtensionClient{
		Client: client,
	}, nil
}
