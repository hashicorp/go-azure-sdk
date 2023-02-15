package apioperationpolicy

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiOperationPolicyClient struct {
	Client  autorest.Client
	baseUri string
}

func NewApiOperationPolicyClientWithBaseURI(endpoint string) ApiOperationPolicyClient {
	return ApiOperationPolicyClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
