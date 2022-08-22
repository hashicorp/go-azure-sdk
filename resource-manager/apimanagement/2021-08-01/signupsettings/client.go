package signupsettings

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignUpSettingsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewSignUpSettingsClientWithBaseURI(endpoint string) SignUpSettingsClient {
	return SignUpSettingsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
