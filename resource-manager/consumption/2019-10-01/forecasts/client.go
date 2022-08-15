package forecasts

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ForecastsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewForecastsClientWithBaseURI(endpoint string) ForecastsClient {
	return ForecastsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
