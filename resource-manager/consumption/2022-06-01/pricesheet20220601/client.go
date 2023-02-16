package pricesheet20220601

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PriceSheet20220601Client struct {
	Client  autorest.Client
	baseUri string
}

func NewPriceSheet20220601ClientWithBaseURI(endpoint string) PriceSheet20220601Client {
	return PriceSheet20220601Client{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
