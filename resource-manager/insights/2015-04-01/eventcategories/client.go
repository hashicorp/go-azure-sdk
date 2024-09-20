package eventcategories

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventCategoriesClient struct {
	Client *resourcemanager.Client
}

func NewEventCategoriesClientWithBaseURI(sdkApi sdkEnv.Api) (*EventCategoriesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "eventcategories", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EventCategoriesClient: %+v", err)
	}

	return &EventCategoriesClient{
		Client: client,
	}, nil
}
