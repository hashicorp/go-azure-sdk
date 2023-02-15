package gatewaycertificateauthority

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GatewayCertificateAuthorityClient struct {
	Client  autorest.Client
	baseUri string
}

func NewGatewayCertificateAuthorityClientWithBaseURI(endpoint string) GatewayCertificateAuthorityClient {
	return GatewayCertificateAuthorityClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
