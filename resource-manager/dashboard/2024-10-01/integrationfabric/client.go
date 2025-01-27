package integrationfabric

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntegrationFabricClient struct {
	Client *resourcemanager.Client
}

func NewIntegrationFabricClientWithBaseURI(sdkApi sdkEnv.Api) (*IntegrationFabricClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "integrationfabric", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IntegrationFabricClient: %+v", err)
	}

	return &IntegrationFabricClient{
		Client: client,
	}, nil
}
