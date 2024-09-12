package entitlementmanagementcatalogcustomworkflowextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementCatalogCustomWorkflowExtensionClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementCatalogCustomWorkflowExtensionClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementCatalogCustomWorkflowExtensionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementcatalogcustomworkflowextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementCatalogCustomWorkflowExtensionClient: %+v", err)
	}

	return &EntitlementManagementCatalogCustomWorkflowExtensionClient{
		Client: client,
	}, nil
}
