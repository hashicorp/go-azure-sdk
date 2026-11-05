package agentsconnectors

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AgentsConnectorsClient struct {
	Client *resourcemanager.Client
}

func NewAgentsConnectorsClientWithBaseURI(sdkApi sdkEnv.Api) (*AgentsConnectorsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "agentsconnectors", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AgentsConnectorsClient: %+v", err)
	}

	return &AgentsConnectorsClient{
		Client: client,
	}, nil
}
