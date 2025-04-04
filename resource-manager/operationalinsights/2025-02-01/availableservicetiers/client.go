package availableservicetiers

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AvailableServiceTiersClient struct {
	Client *resourcemanager.Client
}

func NewAvailableServiceTiersClientWithBaseURI(sdkApi sdkEnv.Api) (*AvailableServiceTiersClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "availableservicetiers", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AvailableServiceTiersClient: %+v", err)
	}

	return &AvailableServiceTiersClient{
		Client: client,
	}, nil
}
