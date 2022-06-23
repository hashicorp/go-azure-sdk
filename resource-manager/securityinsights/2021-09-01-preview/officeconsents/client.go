package officeconsents

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OfficeConsentsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewOfficeConsentsClientWithBaseURI(endpoint string) OfficeConsentsClient {
	return OfficeConsentsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
