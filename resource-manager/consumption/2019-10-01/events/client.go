package events

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewEventsClientWithBaseURI(endpoint string) EventsClient {
	return EventsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
