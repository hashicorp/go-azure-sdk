package containerappsauthconfigs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BlobStorageTokenStore struct {
	BlobContainerUri          *string `json:"blobContainerUri,omitempty"`
	ClientId                  *string `json:"clientId,omitempty"`
	ManagedIdentityResourceId *string `json:"managedIdentityResourceId,omitempty"`
	SasURLSettingName         *string `json:"sasUrlSettingName,omitempty"`
}
