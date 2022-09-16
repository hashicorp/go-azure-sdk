package ipfirewallrules

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IPFirewallRulesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewIPFirewallRulesClientWithBaseURI(endpoint string) IPFirewallRulesClient {
	return IPFirewallRulesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
