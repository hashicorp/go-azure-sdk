package registeredidentities

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegisteredIdentitiesClient struct {
	Client *resourcemanager.Client
}

func NewRegisteredIdentitiesClientWithBaseURI(sdkApi sdkEnv.Api) (*RegisteredIdentitiesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "registeredidentities", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RegisteredIdentitiesClient: %+v", err)
	}

	return &RegisteredIdentitiesClient{
		Client: client,
	}, nil
}
