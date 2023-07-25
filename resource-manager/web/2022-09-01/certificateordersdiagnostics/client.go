package certificateordersdiagnostics

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CertificateOrdersDiagnosticsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewCertificateOrdersDiagnosticsClientWithBaseURI(endpoint string) CertificateOrdersDiagnosticsClient {
	return CertificateOrdersDiagnosticsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
