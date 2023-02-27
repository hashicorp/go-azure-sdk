// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2017_04_01

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
