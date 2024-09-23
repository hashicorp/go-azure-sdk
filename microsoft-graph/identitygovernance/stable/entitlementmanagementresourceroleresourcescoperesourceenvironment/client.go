package entitlementmanagementresourceroleresourcescoperesourceenvironment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementResourceRoleResourceScopeResourceEnvironmentClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementResourceRoleResourceScopeResourceEnvironmentClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementResourceRoleResourceScopeResourceEnvironmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementresourceroleresourcescoperesourceenvironment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementResourceRoleResourceScopeResourceEnvironmentClient: %+v", err)
	}

	return &EntitlementManagementResourceRoleResourceScopeResourceEnvironmentClient{
		Client: client,
	}, nil
}
