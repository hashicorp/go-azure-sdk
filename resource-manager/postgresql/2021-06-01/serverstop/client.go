package serverstop

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerStopClient struct {
	Client *resourcemanager.Client
}

func NewServerStopClientWithBaseURI(api environments.Api) (*ServerStopClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "serverstop", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ServerStopClient: %+v", err)
	}

	return &ServerStopClient{
		Client: client,
	}, nil
}
