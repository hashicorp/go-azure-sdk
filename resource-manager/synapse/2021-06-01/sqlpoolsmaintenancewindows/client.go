package sqlpoolsmaintenancewindows

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolsMaintenanceWindowsClient struct {
	Client *resourcemanager.Client
}

func NewSqlPoolsMaintenanceWindowsClientWithBaseURI(sdkApi sdkEnv.Api) (*SqlPoolsMaintenanceWindowsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "sqlpoolsmaintenancewindows", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SqlPoolsMaintenanceWindowsClient: %+v", err)
	}

	return &SqlPoolsMaintenanceWindowsClient{
		Client: client,
	}, nil
}
