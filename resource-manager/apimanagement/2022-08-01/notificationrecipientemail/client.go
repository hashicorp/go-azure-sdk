package notificationrecipientemail

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NotificationRecipientEmailClient struct {
	Client  autorest.Client
	baseUri string
}

func NewNotificationRecipientEmailClientWithBaseURI(endpoint string) NotificationRecipientEmailClient {
	return NotificationRecipientEmailClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
