package subscriptionraipolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubscriptionRaiPolicyClient struct {
	Client *resourcemanager.Client
}

func NewSubscriptionRaiPolicyClientWithBaseURI(sdkApi sdkEnv.Api) (*SubscriptionRaiPolicyClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "subscriptionraipolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SubscriptionRaiPolicyClient: %+v", err)
	}

	return &SubscriptionRaiPolicyClient{
		Client: client,
	}, nil
}
