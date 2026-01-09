package entitlementmanagementresourcenamespace

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementResourceNamespaceClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementResourceNamespaceClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementResourceNamespaceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementresourcenamespace", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementResourceNamespaceClient: %+v", err)
	}

	return &EntitlementManagementResourceNamespaceClient{
		Client: client,
	}, nil
}
