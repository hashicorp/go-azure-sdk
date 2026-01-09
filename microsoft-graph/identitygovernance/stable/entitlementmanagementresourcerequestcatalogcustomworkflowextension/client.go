package entitlementmanagementresourcerequestcatalogcustomworkflowextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementResourceRequestCatalogCustomWorkflowExtensionClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementResourceRequestCatalogCustomWorkflowExtensionClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementResourceRequestCatalogCustomWorkflowExtensionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementresourcerequestcatalogcustomworkflowextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementResourceRequestCatalogCustomWorkflowExtensionClient: %+v", err)
	}

	return &EntitlementManagementResourceRequestCatalogCustomWorkflowExtensionClient{
		Client: client,
	}, nil
}
