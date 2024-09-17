package location

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LocationClient struct {
	Client *resourcemanager.Client
}

func NewLocationClientWithBaseURI(sdkApi sdkEnv.Api) (*LocationClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "location", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LocationClient: %+v", err)
	}

	return &LocationClient{
		Client: client,
	}, nil
}
