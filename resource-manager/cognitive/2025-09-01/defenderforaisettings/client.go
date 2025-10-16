package defenderforaisettings

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefenderForAISettingsClient struct {
	Client *resourcemanager.Client
}

func NewDefenderForAISettingsClientWithBaseURI(sdkApi sdkEnv.Api) (*DefenderForAISettingsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "defenderforaisettings", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DefenderForAISettingsClient: %+v", err)
	}

	return &DefenderForAISettingsClient{
		Client: client,
	}, nil
}
