package onboardtod4api

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnboardToD4APIClient struct {
	Client *resourcemanager.Client
}

func NewOnboardToD4APIClientWithBaseURI(sdkApi sdkEnv.Api) (*OnboardToD4APIClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "onboardtod4api", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnboardToD4APIClient: %+v", err)
	}

	return &OnboardToD4APIClient{
		Client: client,
	}, nil
}
