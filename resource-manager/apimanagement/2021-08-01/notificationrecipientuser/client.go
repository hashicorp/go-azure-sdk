package notificationrecipientuser

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NotificationRecipientUserClient struct {
	Client  autorest.Client
	baseUri string
}

func NewNotificationRecipientUserClientWithBaseURI(endpoint string) NotificationRecipientUserClient {
	return NotificationRecipientUserClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
