package datalakestoreaccounts

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataLakeStoreAccountsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewDataLakeStoreAccountsClientWithBaseURI(endpoint string) DataLakeStoreAccountsClient {
	return DataLakeStoreAccountsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
