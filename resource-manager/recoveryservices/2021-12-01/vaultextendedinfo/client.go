package vaultextendedinfo

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VaultExtendedInfoClient struct {
	Client  autorest.Client
	baseUri string
}

func NewVaultExtendedInfoClientWithBaseURI(endpoint string) VaultExtendedInfoClient {
	return VaultExtendedInfoClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
