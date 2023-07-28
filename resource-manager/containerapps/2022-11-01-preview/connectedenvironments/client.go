package connectedenvironments

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectedEnvironmentsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewConnectedEnvironmentsClientWithBaseURI(endpoint string) ConnectedEnvironmentsClient {
	return ConnectedEnvironmentsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
