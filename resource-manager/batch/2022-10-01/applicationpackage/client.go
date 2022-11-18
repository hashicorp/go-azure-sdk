package applicationpackage

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationPackageClient struct {
	Client  autorest.Client
	baseUri string
}

func NewApplicationPackageClientWithBaseURI(endpoint string) ApplicationPackageClient {
	return ApplicationPackageClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
