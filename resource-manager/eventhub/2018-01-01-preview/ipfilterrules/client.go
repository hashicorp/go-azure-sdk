package ipfilterrules

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IpFilterRulesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewIpFilterRulesClientWithBaseURI(endpoint string) IpFilterRulesClient {
	return IpFilterRulesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
