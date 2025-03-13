package hunts

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HuntsClient struct {
	Client *resourcemanager.Client
}

func NewHuntsClientWithBaseURI(sdkApi sdkEnv.Api) (*HuntsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "hunts", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating HuntsClient: %+v", err)
	}

	return &HuntsClient{
		Client: client,
	}, nil
}
