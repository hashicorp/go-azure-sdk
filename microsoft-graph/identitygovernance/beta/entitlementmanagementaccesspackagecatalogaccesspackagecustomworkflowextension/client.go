package entitlementmanagementaccesspackagecatalogaccesspackagecustomworkflowextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageCatalogAccessPackageCustomWorkflowExtensionClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageCatalogAccessPackageCustomWorkflowExtensionClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageCatalogAccessPackageCustomWorkflowExtensionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackagecatalogaccesspackagecustomworkflowextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageCatalogAccessPackageCustomWorkflowExtensionClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageCatalogAccessPackageCustomWorkflowExtensionClient{
		Client: client,
	}, nil
}
