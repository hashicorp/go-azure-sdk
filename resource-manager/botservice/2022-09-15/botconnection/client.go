package botconnection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BotConnectionClient struct {
	Client *resourcemanager.Client
}

func NewBotConnectionClientWithBaseURI(sdkApi sdkEnv.Api) (*BotConnectionClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "botconnection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating BotConnectionClient: %+v", err)
	}

	return &BotConnectionClient{
		Client: client,
	}, nil
}
