package graphqlapiresolver

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GraphQLApiResolverClient struct {
	Client  autorest.Client
	baseUri string
}

func NewGraphQLApiResolverClientWithBaseURI(endpoint string) GraphQLApiResolverClient {
	return GraphQLApiResolverClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
