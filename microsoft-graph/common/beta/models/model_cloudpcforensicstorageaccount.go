package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCForensicStorageAccount struct {
	ODataType          *string `json:"@odata.type,omitempty"`
	StorageAccountId   *string `json:"storageAccountId,omitempty"`
	StorageAccountName *string `json:"storageAccountName,omitempty"`
}
