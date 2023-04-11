package advancedthreatprotection

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdvancedThreatProtectionClient struct {
	Client  autorest.Client
	baseUri string
}

func NewAdvancedThreatProtectionClientWithBaseURI(endpoint string) AdvancedThreatProtectionClient {
	return AdvancedThreatProtectionClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
