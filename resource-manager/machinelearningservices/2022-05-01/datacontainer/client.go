package datacontainer

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataContainerClient struct {
	Client  autorest.Client
	baseUri string
}

func NewDataContainerClientWithBaseURI(endpoint string) DataContainerClient {
	return DataContainerClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
