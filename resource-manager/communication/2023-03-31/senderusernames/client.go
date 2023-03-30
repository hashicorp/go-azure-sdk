package senderusernames

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SenderUsernamesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewSenderUsernamesClientWithBaseURI(endpoint string) SenderUsernamesClient {
	return SenderUsernamesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
