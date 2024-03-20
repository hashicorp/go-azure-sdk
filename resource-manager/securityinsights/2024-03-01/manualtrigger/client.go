package manualtrigger

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManualTriggerClient struct {
	Client *resourcemanager.Client
}

func NewManualTriggerClientWithBaseURI(sdkApi sdkEnv.Api) (*ManualTriggerClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "manualtrigger", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManualTriggerClient: %+v", err)
	}

	return &ManualTriggerClient{
		Client: client,
	}, nil
}
