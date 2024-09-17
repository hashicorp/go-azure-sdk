package topology

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TopologyClient struct {
	Client *resourcemanager.Client
}

func NewTopologyClientWithBaseURI(sdkApi sdkEnv.Api) (*TopologyClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "topology", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TopologyClient: %+v", err)
	}

	return &TopologyClient{
		Client: client,
	}, nil
}
