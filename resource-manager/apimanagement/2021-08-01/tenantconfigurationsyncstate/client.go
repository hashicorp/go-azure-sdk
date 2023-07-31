package tenantconfigurationsyncstate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TenantConfigurationSyncStateClient struct {
	Client *resourcemanager.Client
}

func NewTenantConfigurationSyncStateClientWithBaseURI(api sdkEnv.Api) (*TenantConfigurationSyncStateClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "tenantconfigurationsyncstate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TenantConfigurationSyncStateClient: %+v", err)
	}

	return &TenantConfigurationSyncStateClient{
		Client: client,
	}, nil
}
