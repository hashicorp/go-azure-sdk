package checknameavailabilitynamespaces

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CheckNameAvailabilityNamespacesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewCheckNameAvailabilityNamespacesClientWithBaseURI(endpoint string) CheckNameAvailabilityNamespacesClient {
	return CheckNameAvailabilityNamespacesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
