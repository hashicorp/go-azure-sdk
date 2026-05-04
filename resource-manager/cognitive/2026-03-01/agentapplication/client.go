package agentapplication

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AgentApplicationClient struct {
	Client *resourcemanager.Client
}

func NewAgentApplicationClientWithBaseURI(sdkApi sdkEnv.Api) (*AgentApplicationClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "agentapplication", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AgentApplicationClient: %+v", err)
	}

	return &AgentApplicationClient{
		Client: client,
	}, nil
}
