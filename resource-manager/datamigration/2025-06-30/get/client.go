package get

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GETClient struct {
	Client *resourcemanager.Client
}

func NewGETClientWithBaseURI(sdkApi sdkEnv.Api) (*GETClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "get", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GETClient: %+v", err)
	}

	return &GETClient{
		Client: client,
	}, nil
}
