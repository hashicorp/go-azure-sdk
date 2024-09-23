package entitlementmanagement

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagement", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementClient: %+v", err)
	}

	return &EntitlementManagementClient{
		Client: client,
	}, nil
}
