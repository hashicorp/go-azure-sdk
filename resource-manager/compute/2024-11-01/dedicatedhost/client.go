package dedicatedhost

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DedicatedHostClient struct {
	Client *resourcemanager.Client
}

func NewDedicatedHostClientWithBaseURI(sdkApi sdkEnv.Api) (*DedicatedHostClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "dedicatedhost", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DedicatedHostClient: %+v", err)
	}

	return &DedicatedHostClient{
		Client: client,
	}, nil
}
