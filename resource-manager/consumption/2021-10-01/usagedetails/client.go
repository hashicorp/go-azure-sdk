package usagedetails

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UsageDetailsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewUsageDetailsClientWithBaseURI(endpoint string) UsageDetailsClient {
	return UsageDetailsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
