package usersubscription

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserSubscriptionClient struct {
	Client *resourcemanager.Client
}

func NewUserSubscriptionClientWithBaseURI(api environments.Api) (*UserSubscriptionClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "usersubscription", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserSubscriptionClient: %+v", err)
	}

	return &UserSubscriptionClient{
		Client: client,
	}, nil
}
