package integrationruntimenodes

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntegrationRuntimeNodesClient struct {
	Client *resourcemanager.Client
}

func NewIntegrationRuntimeNodesClientWithBaseURI(sdkApi sdkEnv.Api) (*IntegrationRuntimeNodesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "integrationruntimenodes", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IntegrationRuntimeNodesClient: %+v", err)
	}

	return &IntegrationRuntimeNodesClient{
		Client: client,
	}, nil
}
