package hypervhost

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HyperVHostClient struct {
	Client *resourcemanager.Client
}

func NewHyperVHostClientWithBaseURI(sdkApi sdkEnv.Api) (*HyperVHostClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "hypervhost", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating HyperVHostClient: %+v", err)
	}

	return &HyperVHostClient{
		Client: client,
	}, nil
}
