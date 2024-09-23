package entitlementmanagementresourceroleresourcescope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementResourceRoleResourceScopeClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementResourceRoleResourceScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementResourceRoleResourceScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementresourceroleresourcescope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementResourceRoleResourceScopeClient: %+v", err)
	}

	return &EntitlementManagementResourceRoleResourceScopeClient{
		Client: client,
	}, nil
}
