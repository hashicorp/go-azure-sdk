package authorizationaccesspolicy

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthorizationAccessPolicyClient struct {
	Client  autorest.Client
	baseUri string
}

func NewAuthorizationAccessPolicyClientWithBaseURI(endpoint string) AuthorizationAccessPolicyClient {
	return AuthorizationAccessPolicyClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
