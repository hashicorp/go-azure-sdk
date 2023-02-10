package ascusages

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AscUsagesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewAscUsagesClientWithBaseURI(endpoint string) AscUsagesClient {
	return AscUsagesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
