package entitlementmanagementresourcerequestresourcescoperesourceenvironment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementResourceRequestResourceScopeResourceEnvironmentClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementResourceRequestResourceScopeResourceEnvironmentClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementResourceRequestResourceScopeResourceEnvironmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementresourcerequestresourcescoperesourceenvironment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementResourceRequestResourceScopeResourceEnvironmentClient: %+v", err)
	}

	return &EntitlementManagementResourceRequestResourceScopeResourceEnvironmentClient{
		Client: client,
	}, nil
}
