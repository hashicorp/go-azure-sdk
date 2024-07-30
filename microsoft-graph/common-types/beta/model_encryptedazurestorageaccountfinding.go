package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EncryptedAzureStorageAccountFinding struct {
	CreatedDateTime     *string                                                 `json:"createdDateTime,omitempty"`
	EncryptionManagedBy *EncryptedAzureStorageAccountFindingEncryptionManagedBy `json:"encryptionManagedBy,omitempty"`
	Id                  *string                                                 `json:"id,omitempty"`
	ODataType           *string                                                 `json:"@odata.type,omitempty"`
	StorageAccount      *AuthorizationSystemResource                            `json:"storageAccount,omitempty"`
}
