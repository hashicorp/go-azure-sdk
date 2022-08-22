package quotabycounterkeys

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type QuotaByCounterKeysClient struct {
	Client  autorest.Client
	baseUri string
}

func NewQuotaByCounterKeysClientWithBaseURI(endpoint string) QuotaByCounterKeysClient {
	return QuotaByCounterKeysClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
