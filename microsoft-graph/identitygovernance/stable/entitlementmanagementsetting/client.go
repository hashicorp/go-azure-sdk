package entitlementmanagementsetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementSettingClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementSettingClient: %+v", err)
	}

	return &EntitlementManagementSettingClient{
		Client: client,
	}, nil
}
