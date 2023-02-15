package notification

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NotificationClient struct {
	Client  autorest.Client
	baseUri string
}

func NewNotificationClientWithBaseURI(endpoint string) NotificationClient {
	return NotificationClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
