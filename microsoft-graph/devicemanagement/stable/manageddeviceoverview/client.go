package manageddeviceoverview

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceOverviewClient struct {
	Client *msgraph.Client
}

func NewManagedDeviceOverviewClientWithBaseURI(sdkApi sdkEnv.Api) (*ManagedDeviceOverviewClient, error) {
	client, err := msgraph.NewClient(sdkApi, "manageddeviceoverview", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedDeviceOverviewClient: %+v", err)
	}

	return &ManagedDeviceOverviewClient{
		Client: client,
	}, nil
}
