package tenantactivitylogs

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TenantActivityLogsClient struct {
	Client *resourcemanager.Client
}

func NewTenantActivityLogsClientWithBaseURI(sdkApi sdkEnv.Api) (*TenantActivityLogsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "tenantactivitylogs", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TenantActivityLogsClient: %+v", err)
	}

	return &TenantActivityLogsClient{
		Client: client,
	}, nil
}
