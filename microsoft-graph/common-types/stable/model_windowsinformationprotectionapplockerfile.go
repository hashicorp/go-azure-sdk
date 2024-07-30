package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsInformationProtectionAppLockerFile struct {
	DisplayName *string `json:"displayName,omitempty"`
	File        *string `json:"file,omitempty"`
	FileHash    *string `json:"fileHash,omitempty"`
	Id          *string `json:"id,omitempty"`
	ODataType   *string `json:"@odata.type,omitempty"`
	Version     *string `json:"version,omitempty"`
}
