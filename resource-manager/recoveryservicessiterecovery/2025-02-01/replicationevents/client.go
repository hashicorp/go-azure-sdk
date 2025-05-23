package replicationevents

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReplicationEventsClient struct {
	Client *resourcemanager.Client
}

func NewReplicationEventsClientWithBaseURI(sdkApi sdkEnv.Api) (*ReplicationEventsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "replicationevents", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ReplicationEventsClient: %+v", err)
	}

	return &ReplicationEventsClient{
		Client: client,
	}, nil
}
