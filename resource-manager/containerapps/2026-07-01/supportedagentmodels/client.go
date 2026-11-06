package supportedagentmodels

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SupportedAgentModelsClient struct {
	Client *resourcemanager.Client
}

func NewSupportedAgentModelsClientWithBaseURI(sdkApi sdkEnv.Api) (*SupportedAgentModelsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "supportedagentmodels", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SupportedAgentModelsClient: %+v", err)
	}

	return &SupportedAgentModelsClient{
		Client: client,
	}, nil
}
