package recommendedactionsessions

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendedActionSessionsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewRecommendedActionSessionsClientWithBaseURI(endpoint string) RecommendedActionSessionsClient {
	return RecommendedActionSessionsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
