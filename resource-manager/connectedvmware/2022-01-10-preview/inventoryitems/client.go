package inventoryitems

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InventoryItemsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewInventoryItemsClientWithBaseURI(endpoint string) InventoryItemsClient {
	return InventoryItemsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
