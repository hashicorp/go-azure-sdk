package events

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventsClient struct {
	Client *resourcemanager.Client
}

func NewEventsClientWithBaseURI(api environments.Api) (*EventsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "events", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EventsClient: %+v", err)
	}

	return &EventsClient{
		Client: client,
	}, nil
}
