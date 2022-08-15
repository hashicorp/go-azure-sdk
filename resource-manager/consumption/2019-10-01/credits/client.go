package credits

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreditsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewCreditsClientWithBaseURI(endpoint string) CreditsClient {
	return CreditsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
