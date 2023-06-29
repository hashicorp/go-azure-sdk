package defaultaccount

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultAccountClient struct {
	Client  autorest.Client
	baseUri string
}

func NewDefaultAccountClientWithBaseURI(endpoint string) DefaultAccountClient {
	return DefaultAccountClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
