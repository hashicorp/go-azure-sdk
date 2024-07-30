package entitlementmanagementresourceroleresource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementResourceRoleResourceClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementResourceRoleResourceClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementResourceRoleResourceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementresourceroleresource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementResourceRoleResourceClient: %+v", err)
	}

	return &EntitlementManagementResourceRoleResourceClient{
		Client: client,
	}, nil
}
