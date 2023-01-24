package python2package

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Python2PackageClient struct {
	Client  autorest.Client
	baseUri string
}

func NewPython2PackageClientWithBaseURI(endpoint string) Python2PackageClient {
	return Python2PackageClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
