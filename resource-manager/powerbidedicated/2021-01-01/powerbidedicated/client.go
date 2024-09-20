package powerbidedicated

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PowerBIDedicatedClient struct {
	Client *resourcemanager.Client
}

func NewPowerBIDedicatedClientWithBaseURI(sdkApi sdkEnv.Api) (*PowerBIDedicatedClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "powerbidedicated", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PowerBIDedicatedClient: %+v", err)
	}

	return &PowerBIDedicatedClient{
		Client: client,
	}, nil
}
