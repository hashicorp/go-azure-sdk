package securescorecontroldefinitions

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecureScoreControlDefinitionsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewSecureScoreControlDefinitionsClientWithBaseURI(endpoint string) SecureScoreControlDefinitionsClient {
	return SecureScoreControlDefinitionsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
