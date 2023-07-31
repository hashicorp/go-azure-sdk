package dashboard

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DashboardClient struct {
	Client *resourcemanager.Client
}

func NewDashboardClientWithBaseURI(api sdkEnv.Api) (*DashboardClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "dashboard", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DashboardClient: %+v", err)
	}

	return &DashboardClient{
		Client: client,
	}, nil
}
