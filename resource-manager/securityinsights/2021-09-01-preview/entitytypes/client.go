package entitytypes

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntityTypesClient struct {
	Client *resourcemanager.Client
}

func NewEntityTypesClientWithBaseURI(sdkApi sdkEnv.Api) (*EntityTypesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "entitytypes", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntityTypesClient: %+v", err)
	}

	return &EntityTypesClient{
		Client: client,
	}, nil
}
