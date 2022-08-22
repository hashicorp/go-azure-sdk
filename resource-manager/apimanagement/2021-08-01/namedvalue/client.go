package namedvalue

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NamedValueClient struct {
	Client  autorest.Client
	baseUri string
}

func NewNamedValueClientWithBaseURI(endpoint string) NamedValueClient {
	return NamedValueClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
