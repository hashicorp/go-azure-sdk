package incrementalrestorepoints

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IncrementalRestorePointsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewIncrementalRestorePointsClientWithBaseURI(endpoint string) IncrementalRestorePointsClient {
	return IncrementalRestorePointsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
