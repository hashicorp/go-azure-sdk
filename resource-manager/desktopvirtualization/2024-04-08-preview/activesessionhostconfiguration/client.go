package activesessionhostconfiguration

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActiveSessionHostConfigurationClient struct {
	Client *resourcemanager.Client
}

func NewActiveSessionHostConfigurationClientWithBaseURI(sdkApi sdkEnv.Api) (*ActiveSessionHostConfigurationClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "activesessionhostconfiguration", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ActiveSessionHostConfigurationClient: %+v", err)
	}

	return &ActiveSessionHostConfigurationClient{
		Client: client,
	}, nil
}
