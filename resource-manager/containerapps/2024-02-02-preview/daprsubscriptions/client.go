package daprsubscriptions

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DaprSubscriptionsClient struct {
	Client *resourcemanager.Client
}

func NewDaprSubscriptionsClientWithBaseURI(sdkApi sdkEnv.Api) (*DaprSubscriptionsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "daprsubscriptions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DaprSubscriptionsClient: %+v", err)
	}

	return &DaprSubscriptionsClient{
		Client: client,
	}, nil
}
