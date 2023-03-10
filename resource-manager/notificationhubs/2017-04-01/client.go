package v2017_04_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/notificationhubs/2017-04-01/namespaces"
	"github.com/hashicorp/go-azure-sdk/resource-manager/notificationhubs/2017-04-01/notificationhubs"
)

type Client struct {
	Namespaces       *namespaces.NamespacesClient
	NotificationHubs *notificationhubs.NotificationHubsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	namespacesClient := namespaces.NewNamespacesClientWithBaseURI(endpoint)
	configureAuthFunc(&namespacesClient.Client)

	notificationHubsClient := notificationhubs.NewNotificationHubsClientWithBaseURI(endpoint)
	configureAuthFunc(&notificationHubsClient.Client)

	return Client{
		Namespaces:       &namespacesClient,
		NotificationHubs: &notificationHubsClient,
	}
}
