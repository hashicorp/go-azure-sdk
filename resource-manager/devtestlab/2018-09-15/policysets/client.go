package policysets

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicySetsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewPolicySetsClientWithBaseURI(endpoint string) PolicySetsClient {
	return PolicySetsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
