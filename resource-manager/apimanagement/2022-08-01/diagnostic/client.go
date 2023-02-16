package diagnostic

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiagnosticClient struct {
	Client  autorest.Client
	baseUri string
}

func NewDiagnosticClientWithBaseURI(endpoint string) DiagnosticClient {
	return DiagnosticClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
