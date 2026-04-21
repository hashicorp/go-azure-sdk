package serverconfigurationoptions

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerConfigurationOptionsClient struct {
	Client *resourcemanager.Client
}

func NewServerConfigurationOptionsClientWithBaseURI(sdkApi sdkEnv.Api) (*ServerConfigurationOptionsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "serverconfigurationoptions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ServerConfigurationOptionsClient: %+v", err)
	}

	return &ServerConfigurationOptionsClient{
		Client: client,
	}, nil
}
