package automationrule

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutomationRuleClient struct {
	Client  autorest.Client
	baseUri string
}

func NewAutomationRuleClientWithBaseURI(endpoint string) AutomationRuleClient {
	return AutomationRuleClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
