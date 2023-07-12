package v2017_04_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/notificationhubs/2017-04-01/namespaces"
	"github.com/hashicorp/go-azure-sdk/resource-manager/notificationhubs/2017-04-01/notificationhubs"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Namespaces       *namespaces.NamespacesClient
	NotificationHubs *notificationhubs.NotificationHubsClient
}

func NewClientWithBaseURI(api environments.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	namespacesClient, err := namespaces.NewNamespacesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Namespaces client: %+v", err)
	}
	configureFunc(namespacesClient.Client)

	notificationHubsClient, err := notificationhubs.NewNotificationHubsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building NotificationHubs client: %+v", err)
	}
	configureFunc(notificationHubsClient.Client)

	return &Client{
		Namespaces:       namespacesClient,
		NotificationHubs: notificationHubsClient,
	}, nil
}
