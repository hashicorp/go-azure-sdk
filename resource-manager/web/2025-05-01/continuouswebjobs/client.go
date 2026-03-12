package continuouswebjobs

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContinuousWebJobsClient struct {
	Client *resourcemanager.Client
}

func NewContinuousWebJobsClientWithBaseURI(sdkApi sdkEnv.Api) (*ContinuousWebJobsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "continuouswebjobs", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ContinuousWebJobsClient: %+v", err)
	}

	return &ContinuousWebJobsClient{
		Client: client,
	}, nil
}
