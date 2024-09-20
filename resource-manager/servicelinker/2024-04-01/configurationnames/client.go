package configurationnames

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigurationNamesClient struct {
	Client *resourcemanager.Client
}

func NewConfigurationNamesClientWithBaseURI(sdkApi sdkEnv.Api) (*ConfigurationNamesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "configurationnames", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ConfigurationNamesClient: %+v", err)
	}

	return &ConfigurationNamesClient{
		Client: client,
	}, nil
}
