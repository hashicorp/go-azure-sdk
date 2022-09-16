package restorabledroppedsqlpools

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestorableDroppedSqlPoolsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewRestorableDroppedSqlPoolsClientWithBaseURI(endpoint string) RestorableDroppedSqlPoolsClient {
	return RestorableDroppedSqlPoolsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
