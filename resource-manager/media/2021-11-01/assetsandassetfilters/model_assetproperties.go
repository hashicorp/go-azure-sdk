package assetsandassetfilters

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssetProperties struct {
	AlternateId             *string                       `json:"alternateId,omitempty"`
	AssetId                 *string                       `json:"assetId,omitempty"`
	Container               *string                       `json:"container,omitempty"`
	Created                 *string                       `json:"created,omitempty"`
	Description             *string                       `json:"description,omitempty"`
	LastModified            *string                       `json:"lastModified,omitempty"`
	StorageAccountName      *string                       `json:"storageAccountName,omitempty"`
	StorageEncryptionFormat *AssetStorageEncryptionFormat `json:"storageEncryptionFormat,omitempty"`
}
