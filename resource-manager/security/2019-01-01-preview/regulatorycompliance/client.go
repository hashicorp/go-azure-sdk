package regulatorycompliance

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegulatoryComplianceClient struct {
	Client  autorest.Client
	baseUri string
}

func NewRegulatoryComplianceClientWithBaseURI(endpoint string) RegulatoryComplianceClient {
	return RegulatoryComplianceClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
