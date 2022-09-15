package restorepointcollections

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestorePointCollectionsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewRestorePointCollectionsClientWithBaseURI(endpoint string) RestorePointCollectionsClient {
	return RestorePointCollectionsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
