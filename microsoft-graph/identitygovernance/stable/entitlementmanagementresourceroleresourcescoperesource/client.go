package entitlementmanagementresourceroleresourcescoperesource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementResourceRoleResourceScopeResourceClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementResourceRoleResourceScopeResourceClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementResourceRoleResourceScopeResourceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementresourceroleresourcescoperesource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementResourceRoleResourceScopeResourceClient: %+v", err)
	}

	return &EntitlementManagementResourceRoleResourceScopeResourceClient{
		Client: client,
	}, nil
}
