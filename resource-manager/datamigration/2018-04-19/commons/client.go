package commons

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommonsClient struct {
	Client *resourcemanager.Client
}

func NewCommonsClientWithBaseURI(sdkApi sdkEnv.Api) (*CommonsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "commons", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CommonsClient: %+v", err)
	}

	return &CommonsClient{
		Client: client,
	}, nil
}
