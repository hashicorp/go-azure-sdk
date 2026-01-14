package entities

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitiesClient struct {
	Client *resourcemanager.Client
}

func NewEntitiesClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitiesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "entities", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitiesClient: %+v", err)
	}

	return &EntitiesClient{
		Client: client,
	}, nil
}
