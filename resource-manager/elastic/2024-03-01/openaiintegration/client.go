package openaiintegration

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OpenAIIntegrationClient struct {
	Client *resourcemanager.Client
}

func NewOpenAIIntegrationClientWithBaseURI(sdkApi sdkEnv.Api) (*OpenAIIntegrationClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "openaiintegration", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OpenAIIntegrationClient: %+v", err)
	}

	return &OpenAIIntegrationClient{
		Client: client,
	}, nil
}
