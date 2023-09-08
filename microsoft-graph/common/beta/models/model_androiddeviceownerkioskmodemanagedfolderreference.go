package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerKioskModeManagedFolderReference struct {
	FolderIdentifier *string `json:"folderIdentifier,omitempty"`
	FolderName       *string `json:"folderName,omitempty"`
	ODataType        *string `json:"@odata.type,omitempty"`
}
