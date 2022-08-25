package serverupgrade

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerUpgradeClient struct {
	Client  autorest.Client
	baseUri string
}

func NewServerUpgradeClientWithBaseURI(endpoint string) ServerUpgradeClient {
	return ServerUpgradeClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
