package entitlementmanagementresourceenvironmentresourcescoperesourceroleresourceenvironment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementResourceEnvironmentResourceScopeResourceRoleResourceEnvironmentClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementResourceEnvironmentResourceScopeResourceRoleResourceEnvironmentClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementResourceEnvironmentResourceScopeResourceRoleResourceEnvironmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementresourceenvironmentresourcescoperesourceroleresourceenvironment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementResourceEnvironmentResourceScopeResourceRoleResourceEnvironmentClient: %+v", err)
	}

	return &EntitlementManagementResourceEnvironmentResourceScopeResourceRoleResourceEnvironmentClient{
		Client: client,
	}, nil
}
