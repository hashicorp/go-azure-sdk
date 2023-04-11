package securescore

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecureScoreClient struct {
	Client  autorest.Client
	baseUri string
}

func NewSecureScoreClientWithBaseURI(endpoint string) SecureScoreClient {
	return SecureScoreClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
