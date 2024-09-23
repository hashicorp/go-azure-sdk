package entitlementmanagementresourcerolescopescoperesourcescope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementResourceRoleScopeScopeResourceScopeClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementResourceRoleScopeScopeResourceScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementResourceRoleScopeScopeResourceScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementresourcerolescopescoperesourcescope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementResourceRoleScopeScopeResourceScopeClient: %+v", err)
	}

	return &EntitlementManagementResourceRoleScopeScopeResourceScopeClient{
		Client: client,
	}, nil
}
