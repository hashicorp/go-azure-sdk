package storageaccounts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EncryptionService struct {
	Enabled         *bool    `json:"enabled,omitempty"`
	KeyType         *KeyType `json:"keyType,omitempty"`
	LastEnabledTime *string  `json:"lastEnabledTime,omitempty"`
}
