package entitlementmanagementresourcerolescope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementResourceRoleScopeClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementResourceRoleScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementResourceRoleScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementresourcerolescope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementResourceRoleScopeClient: %+v", err)
	}

	return &EntitlementManagementResourceRoleScopeClient{
		Client: client,
	}, nil
}
