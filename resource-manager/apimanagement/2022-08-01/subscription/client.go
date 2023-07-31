package subscription

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubscriptionClient struct {
	Client *resourcemanager.Client
}

func NewSubscriptionClientWithBaseURI(api sdkEnv.Api) (*SubscriptionClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "subscription", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SubscriptionClient: %+v", err)
	}

	return &SubscriptionClient{
		Client: client,
	}, nil
}
