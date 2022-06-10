package regions

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegionsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewRegionsClientWithBaseURI(endpoint string) RegionsClient {
	return RegionsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
