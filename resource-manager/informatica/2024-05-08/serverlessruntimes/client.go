package serverlessruntimes

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerlessRuntimesClient struct {
	Client *resourcemanager.Client
}

func NewServerlessRuntimesClientWithBaseURI(sdkApi sdkEnv.Api) (*ServerlessRuntimesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "serverlessruntimes", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ServerlessRuntimesClient: %+v", err)
	}

	return &ServerlessRuntimesClient{
		Client: client,
	}, nil
}
