package openaiintegrationrpmodels

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OpenAIIntegrationRPModelsClient struct {
	Client *resourcemanager.Client
}

func NewOpenAIIntegrationRPModelsClientWithBaseURI(sdkApi sdkEnv.Api) (*OpenAIIntegrationRPModelsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "openaiintegrationrpmodels", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OpenAIIntegrationRPModelsClient: %+v", err)
	}

	return &OpenAIIntegrationRPModelsClient{
		Client: client,
	}, nil
}
