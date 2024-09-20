package metrics

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MetricsClient struct {
	Client *resourcemanager.Client
}

func NewMetricsClientWithBaseURI(sdkApi sdkEnv.Api) (*MetricsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "metrics", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MetricsClient: %+v", err)
	}

	return &MetricsClient{
		Client: client,
	}, nil
}
