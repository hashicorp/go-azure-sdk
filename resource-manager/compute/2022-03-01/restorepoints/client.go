package restorepoints

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestorePointsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewRestorePointsClientWithBaseURI(endpoint string) RestorePointsClient {
	return RestorePointsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
