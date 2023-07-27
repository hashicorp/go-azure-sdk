package performconnectivitycheck

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PerformConnectivityCheckClient struct {
	Client  autorest.Client
	baseUri string
}

func NewPerformConnectivityCheckClientWithBaseURI(endpoint string) PerformConnectivityCheckClient {
	return PerformConnectivityCheckClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
