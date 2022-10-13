package guestagents

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GuestAgentsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewGuestAgentsClientWithBaseURI(endpoint string) GuestAgentsClient {
	return GuestAgentsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
