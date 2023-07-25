package diagnostics

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiagnosticsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewDiagnosticsClientWithBaseURI(endpoint string) DiagnosticsClient {
	return DiagnosticsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
