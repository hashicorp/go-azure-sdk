package alertrulesapis

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlertRulesAPIsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewAlertRulesAPIsClientWithBaseURI(endpoint string) AlertRulesAPIsClient {
	return AlertRulesAPIsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
