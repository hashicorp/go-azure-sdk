package apidiagnostic

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiDiagnosticClient struct {
	Client  autorest.Client
	baseUri string
}

func NewApiDiagnosticClientWithBaseURI(endpoint string) ApiDiagnosticClient {
	return ApiDiagnosticClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
