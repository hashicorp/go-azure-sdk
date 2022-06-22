package msixpackage

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MSIXPackageClient struct {
	Client  autorest.Client
	baseUri string
}

func NewMSIXPackageClientWithBaseURI(endpoint string) MSIXPackageClient {
	return MSIXPackageClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
