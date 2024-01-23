package bandwidthschedules

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BandwidthSchedulesClient struct {
	Client *resourcemanager.Client
}

func NewBandwidthSchedulesClientWithBaseURI(sdkApi sdkEnv.Api) (*BandwidthSchedulesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "bandwidthschedules", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating BandwidthSchedulesClient: %+v", err)
	}

	return &BandwidthSchedulesClient{
		Client: client,
	}, nil
}
