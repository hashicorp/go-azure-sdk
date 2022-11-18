package python3package

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Python3PackageClient struct {
	Client  autorest.Client
	baseUri string
}

func NewPython3PackageClientWithBaseURI(endpoint string) Python3PackageClient {
	return Python3PackageClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
