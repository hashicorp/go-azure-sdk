package compliances

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CompliancesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewCompliancesClientWithBaseURI(endpoint string) CompliancesClient {
	return CompliancesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
