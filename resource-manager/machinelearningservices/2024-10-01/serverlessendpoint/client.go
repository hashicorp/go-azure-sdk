package serverlessendpoint

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerlessEndpointClient struct {
	Client *resourcemanager.Client
}

func NewServerlessEndpointClientWithBaseURI(sdkApi sdkEnv.Api) (*ServerlessEndpointClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "serverlessendpoint", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ServerlessEndpointClient: %+v", err)
	}

	return &ServerlessEndpointClient{
		Client: client,
	}, nil
}
