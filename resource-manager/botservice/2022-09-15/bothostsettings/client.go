package bothostsettings

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BotHostSettingsClient struct {
	Client *resourcemanager.Client
}

func NewBotHostSettingsClientWithBaseURI(sdkApi sdkEnv.Api) (*BotHostSettingsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "bothostsettings", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating BotHostSettingsClient: %+v", err)
	}

	return &BotHostSettingsClient{
		Client: client,
	}, nil
}
