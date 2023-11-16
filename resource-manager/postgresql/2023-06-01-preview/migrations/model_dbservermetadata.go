package migrations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DbServerMetadata struct {
	Location  *string    `json:"location,omitempty"`
	Sku       *ServerSku `json:"sku,omitempty"`
	StorageMb *int64     `json:"storageMb,omitempty"`
	Version   *string    `json:"version,omitempty"`
}
