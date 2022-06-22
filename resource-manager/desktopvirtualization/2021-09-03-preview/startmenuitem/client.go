package startmenuitem

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StartMenuItemClient struct {
	Client  autorest.Client
	baseUri string
}

func NewStartMenuItemClientWithBaseURI(endpoint string) StartMenuItemClient {
	return StartMenuItemClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
