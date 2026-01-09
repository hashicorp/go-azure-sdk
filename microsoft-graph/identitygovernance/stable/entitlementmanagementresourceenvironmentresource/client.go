package entitlementmanagementresourceenvironmentresource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementResourceEnvironmentResourceClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementResourceEnvironmentResourceClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementResourceEnvironmentResourceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementresourceenvironmentresource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementResourceEnvironmentResourceClient: %+v", err)
	}

	return &EntitlementManagementResourceEnvironmentResourceClient{
		Client: client,
	}, nil
}
