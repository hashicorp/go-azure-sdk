package relays

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RelaysClient struct {
	Client *resourcemanager.Client
}

func NewRelaysClientWithBaseURI(sdkApi sdkEnv.Api) (*RelaysClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "relays", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RelaysClient: %+v", err)
	}

	return &RelaysClient{
		Client: client,
	}, nil
}
