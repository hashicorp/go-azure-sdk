package forecast

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ForecastClient struct {
	Client *resourcemanager.Client
}

func NewForecastClientWithBaseURI(api environments.Api) (*ForecastClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "forecast", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ForecastClient: %+v", err)
	}

	return &ForecastClient{
		Client: client,
	}, nil
}
