package prediction

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PredictionClient struct {
	Client  autorest.Client
	baseUri string
}

func NewPredictionClientWithBaseURI(endpoint string) PredictionClient {
	return PredictionClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
