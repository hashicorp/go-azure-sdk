package manageddatabasesecurityevents

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDatabaseSecurityEventsClient struct {
	Client *resourcemanager.Client
}

func NewManagedDatabaseSecurityEventsClientWithBaseURI(sdkApi sdkEnv.Api) (*ManagedDatabaseSecurityEventsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "manageddatabasesecurityevents", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedDatabaseSecurityEventsClient: %+v", err)
	}

	return &ManagedDatabaseSecurityEventsClient{
		Client: client,
	}, nil
}
