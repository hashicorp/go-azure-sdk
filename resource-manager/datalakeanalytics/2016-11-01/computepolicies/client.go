package computepolicies

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComputePoliciesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewComputePoliciesClientWithBaseURI(endpoint string) ComputePoliciesClient {
	return ComputePoliciesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
