package offboardfromd4api

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OffboardFromD4APIClient struct {
	Client *resourcemanager.Client
}

func NewOffboardFromD4APIClientWithBaseURI(sdkApi sdkEnv.Api) (*OffboardFromD4APIClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "offboardfromd4api", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OffboardFromD4APIClient: %+v", err)
	}

	return &OffboardFromD4APIClient{
		Client: client,
	}, nil
}
