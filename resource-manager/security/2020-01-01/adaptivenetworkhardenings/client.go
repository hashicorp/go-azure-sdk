package adaptivenetworkhardenings

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdaptiveNetworkHardeningsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewAdaptiveNetworkHardeningsClientWithBaseURI(endpoint string) AdaptiveNetworkHardeningsClient {
	return AdaptiveNetworkHardeningsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
