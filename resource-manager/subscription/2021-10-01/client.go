package v2021_10_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/subscription/2021-10-01/subscriptions"
)

type Client struct {
	Subscriptions *subscriptions.SubscriptionsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	subscriptionsClient := subscriptions.NewSubscriptionsClientWithBaseURI(endpoint)
	configureAuthFunc(&subscriptionsClient.Client)

	return Client{
		Subscriptions: &subscriptionsClient,
	}
}
