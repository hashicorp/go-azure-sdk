package v2022_12_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2022-12-01-preview/defenderforstorage"
)

type Client struct {
	DefenderForStorage *defenderforstorage.DefenderForStorageClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	defenderForStorageClient := defenderforstorage.NewDefenderForStorageClientWithBaseURI(endpoint)
	configureAuthFunc(&defenderForStorageClient.Client)

	return Client{
		DefenderForStorage: &defenderForStorageClient,
	}
}
