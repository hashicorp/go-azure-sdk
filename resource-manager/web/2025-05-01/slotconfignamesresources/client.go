package slotconfignamesresources

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SlotConfigNamesResourcesClient struct {
	Client *resourcemanager.Client
}

func NewSlotConfigNamesResourcesClientWithBaseURI(sdkApi sdkEnv.Api) (*SlotConfigNamesResourcesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "slotconfignamesresources", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SlotConfigNamesResourcesClient: %+v", err)
	}

	return &SlotConfigNamesResourcesClient{
		Client: client,
	}, nil
}
