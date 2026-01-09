package entitlementmanagementresource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementResourceClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementResourceClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementResourceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementresource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementResourceClient: %+v", err)
	}

	return &EntitlementManagementResourceClient{
		Client: client,
	}, nil
}
