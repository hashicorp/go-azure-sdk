package appresiliency

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppResiliencyClient struct {
	Client *resourcemanager.Client
}

func NewAppResiliencyClientWithBaseURI(sdkApi sdkEnv.Api) (*AppResiliencyClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "appresiliency", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AppResiliencyClient: %+v", err)
	}

	return &AppResiliencyClient{
		Client: client,
	}, nil
}
