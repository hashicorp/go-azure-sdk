package entitlementmanagementresourcescoperesourceroleresource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementResourceScopeResourceRoleResourceClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementResourceScopeResourceRoleResourceClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementResourceScopeResourceRoleResourceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementresourcescoperesourceroleresource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementResourceScopeResourceRoleResourceClient: %+v", err)
	}

	return &EntitlementManagementResourceScopeResourceRoleResourceClient{
		Client: client,
	}, nil
}
