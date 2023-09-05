package configurationprofilesversions

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigurationProfilesVersionsClient struct {
	Client *resourcemanager.Client
}

func NewConfigurationProfilesVersionsClientWithBaseURI(sdkApi sdkEnv.Api) (*ConfigurationProfilesVersionsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "configurationprofilesversions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ConfigurationProfilesVersionsClient: %+v", err)
	}

	return &ConfigurationProfilesVersionsClient{
		Client: client,
	}, nil
}
