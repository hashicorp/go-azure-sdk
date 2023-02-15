package apipolicy

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiPolicyClient struct {
	Client  autorest.Client
	baseUri string
}

func NewApiPolicyClientWithBaseURI(endpoint string) ApiPolicyClient {
	return ApiPolicyClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
