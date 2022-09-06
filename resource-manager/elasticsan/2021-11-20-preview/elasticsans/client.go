package elasticsans

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ElasticSansClient struct {
	Client  autorest.Client
	baseUri string
}

func NewElasticSansClientWithBaseURI(endpoint string) ElasticSansClient {
	return ElasticSansClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
