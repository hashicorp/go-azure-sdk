package alertssuppressionrules

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlertsSuppressionRulesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewAlertsSuppressionRulesClientWithBaseURI(endpoint string) AlertsSuppressionRulesClient {
	return AlertsSuppressionRulesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
