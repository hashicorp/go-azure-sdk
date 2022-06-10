package vmcollectionupdate

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VMCollectionUpdateClient struct {
	Client  autorest.Client
	baseUri string
}

func NewVMCollectionUpdateClientWithBaseURI(endpoint string) VMCollectionUpdateClient {
	return VMCollectionUpdateClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
