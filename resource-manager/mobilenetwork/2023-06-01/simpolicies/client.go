package simpolicies

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SIMPoliciesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewSIMPoliciesClientWithBaseURI(endpoint string) SIMPoliciesClient {
	return SIMPoliciesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
