package databasemigrations

import (
	"github.com/hashicorp/go-azure-helpers/resourcemanager/identity"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureBlob struct {
	AccountKey               *string                                  `json:"accountKey,omitempty"`
	AuthType                 *AuthType                                `json:"authType,omitempty"`
	BlobContainerName        *string                                  `json:"blobContainerName,omitempty"`
	Identity                 *identity.LegacySystemAndUserAssignedMap `json:"identity,omitempty"`
	StorageAccountResourceId *string                                  `json:"storageAccountResourceId,omitempty"`
}
