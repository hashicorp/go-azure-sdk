package modelcapacities

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ModelCapacitiesClient struct {
	Client *resourcemanager.Client
}

func NewModelCapacitiesClientWithBaseURI(sdkApi sdkEnv.Api) (*ModelCapacitiesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "modelcapacities", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ModelCapacitiesClient: %+v", err)
	}

	return &ModelCapacitiesClient{
		Client: client,
	}, nil
}
