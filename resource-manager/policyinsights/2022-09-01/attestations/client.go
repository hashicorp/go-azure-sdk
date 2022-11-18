package attestations

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttestationsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewAttestationsClientWithBaseURI(endpoint string) AttestationsClient {
	return AttestationsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
