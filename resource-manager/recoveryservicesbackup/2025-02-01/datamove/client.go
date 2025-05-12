package datamove

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataMoveClient struct {
	Client  autorest.Client
	baseUri string
}

func NewDataMoveClientWithBaseURI(endpoint string) DataMoveClient {
	return DataMoveClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
