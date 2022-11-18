package batchmanagements

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BatchManagementsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewBatchManagementsClientWithBaseURI(endpoint string) BatchManagementsClient {
	return BatchManagementsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
