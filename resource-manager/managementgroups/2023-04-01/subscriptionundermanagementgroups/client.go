package subscriptionundermanagementgroups

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubscriptionUnderManagementGroupsClient struct {
	Client *resourcemanager.Client
}

func NewSubscriptionUnderManagementGroupsClientWithBaseURI(sdkApi sdkEnv.Api) (*SubscriptionUnderManagementGroupsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "subscriptionundermanagementgroups", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SubscriptionUnderManagementGroupsClient: %+v", err)
	}

	return &SubscriptionUnderManagementGroupsClient{
		Client: client,
	}, nil
}
