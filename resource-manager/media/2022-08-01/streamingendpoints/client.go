package streamingendpoints

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StreamingEndpointsClient struct {
	Client *resourcemanager.Client
}

func NewStreamingEndpointsClientWithBaseURI(sdkApi sdkEnv.Api) (*StreamingEndpointsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "streamingendpoints", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating StreamingEndpointsClient: %+v", err)
	}

	return &StreamingEndpointsClient{
		Client: client,
	}, nil
}
