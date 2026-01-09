package entitlementmanagementresourcerequest

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementResourceRequestClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementResourceRequestClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementResourceRequestClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementresourcerequest", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementResourceRequestClient: %+v", err)
	}

	return &EntitlementManagementResourceRequestClient{
		Client: client,
	}, nil
}
