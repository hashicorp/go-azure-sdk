package supportedoperatingsystems

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SupportedOperatingSystemsClient struct {
	Client *resourcemanager.Client
}

func NewSupportedOperatingSystemsClientWithBaseURI(sdkApi sdkEnv.Api) (*SupportedOperatingSystemsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "supportedoperatingsystems", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SupportedOperatingSystemsClient: %+v", err)
	}

	return &SupportedOperatingSystemsClient{
		Client: client,
	}, nil
}
