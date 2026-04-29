package autoscales

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutoScalesClient struct {
	Client *resourcemanager.Client
}

func NewAutoScalesClientWithBaseURI(sdkApi sdkEnv.Api) (*AutoScalesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "autoscales", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AutoScalesClient: %+v", err)
	}

	return &AutoScalesClient{
		Client: client,
	}, nil
}
