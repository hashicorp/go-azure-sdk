package put

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PUTClient struct {
	Client *resourcemanager.Client
}

func NewPUTClientWithBaseURI(api sdkEnv.Api) (*PUTClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "put", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PUTClient: %+v", err)
	}

	return &PUTClient{
		Client: client,
	}, nil
}
