package resetqueryperformanceinsightdata

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResetQueryPerformanceInsightDataClient struct {
	Client *resourcemanager.Client
}

func NewResetQueryPerformanceInsightDataClientWithBaseURI(sdkApi sdkEnv.Api) (*ResetQueryPerformanceInsightDataClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "resetqueryperformanceinsightdata", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ResetQueryPerformanceInsightDataClient: %+v", err)
	}

	return &ResetQueryPerformanceInsightDataClient{
		Client: client,
	}, nil
}
