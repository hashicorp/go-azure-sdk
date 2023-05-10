package v2022_12_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2022-12-01/subscriptions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2022-12-01/tenants"
)

type Client struct {
	Subscriptions *subscriptions.SubscriptionsClient
	Tenants       *tenants.TenantsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	subscriptionsClient := subscriptions.NewSubscriptionsClientWithBaseURI(endpoint)
	configureAuthFunc(&subscriptionsClient.Client)

	tenantsClient := tenants.NewTenantsClientWithBaseURI(endpoint)
	configureAuthFunc(&tenantsClient.Client)

	return Client{
		Subscriptions: &subscriptionsClient,
		Tenants:       &tenantsClient,
	}
}
