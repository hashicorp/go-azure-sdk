package agents

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AgentsClient struct {
	Client *resourcemanager.Client
}

func NewAgentsClientWithBaseURI(api sdkEnv.Api) (*AgentsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "agents", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AgentsClient: %+v", err)
	}

	return &AgentsClient{
		Client: client,
	}, nil
}
