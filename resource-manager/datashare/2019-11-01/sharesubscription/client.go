package sharesubscription

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ShareSubscriptionClient struct {
	Client *resourcemanager.Client
}

func NewShareSubscriptionClientWithBaseURI(api environments.Api) (*ShareSubscriptionClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "sharesubscription", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ShareSubscriptionClient: %+v", err)
	}

	return &ShareSubscriptionClient{
		Client: client,
	}, nil
}
