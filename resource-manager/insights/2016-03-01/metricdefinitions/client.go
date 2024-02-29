package metricdefinitions

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MetricDefinitionsClient struct {
	Client *resourcemanager.Client
}

func NewMetricDefinitionsClientWithBaseURI(sdkApi sdkEnv.Api) (*MetricDefinitionsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "metricdefinitions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MetricDefinitionsClient: %+v", err)
	}

	return &MetricDefinitionsClient{
		Client: client,
	}, nil
}
