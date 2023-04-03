package alertruleincidents

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlertRuleIncidentsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewAlertRuleIncidentsClientWithBaseURI(endpoint string) AlertRuleIncidentsClient {
	return AlertRuleIncidentsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
