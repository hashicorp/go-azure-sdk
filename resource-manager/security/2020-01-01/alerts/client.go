package alerts

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlertsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewAlertsClientWithBaseURI(endpoint string) AlertsClient {
	return AlertsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
