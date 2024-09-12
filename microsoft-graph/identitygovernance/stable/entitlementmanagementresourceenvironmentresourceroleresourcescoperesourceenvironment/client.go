package entitlementmanagementresourceenvironmentresourceroleresourcescoperesourceenvironment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementResourceEnvironmentResourceRoleResourceScopeResourceEnvironmentClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementResourceEnvironmentResourceRoleResourceScopeResourceEnvironmentClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementResourceEnvironmentResourceRoleResourceScopeResourceEnvironmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementresourceenvironmentresourceroleresourcescoperesourceenvironment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementResourceEnvironmentResourceRoleResourceScopeResourceEnvironmentClient: %+v", err)
	}

	return &EntitlementManagementResourceEnvironmentResourceRoleResourceScopeResourceEnvironmentClient{
		Client: client,
	}, nil
}
