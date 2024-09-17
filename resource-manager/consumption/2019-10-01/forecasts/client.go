package forecasts

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ForecastsClient struct {
	Client *resourcemanager.Client
}

func NewForecastsClientWithBaseURI(sdkApi sdkEnv.Api) (*ForecastsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "forecasts", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ForecastsClient: %+v", err)
	}

	return &ForecastsClient{
		Client: client,
	}, nil
}
