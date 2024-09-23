package entitlementmanagementresourcescoperesourceenvironment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementResourceScopeResourceEnvironmentClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementResourceScopeResourceEnvironmentClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementResourceScopeResourceEnvironmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementresourcescoperesourceenvironment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementResourceScopeResourceEnvironmentClient: %+v", err)
	}

	return &EntitlementManagementResourceScopeResourceEnvironmentClient{
		Client: client,
	}, nil
}
