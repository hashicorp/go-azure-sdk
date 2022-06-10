package vmingestiondetails

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VMIngestionDetailsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewVMIngestionDetailsClientWithBaseURI(endpoint string) VMIngestionDetailsClient {
	return VMIngestionDetailsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
