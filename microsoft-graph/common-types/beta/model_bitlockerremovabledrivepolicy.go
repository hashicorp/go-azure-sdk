package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BitLockerRemovableDrivePolicy struct {
	BlockCrossOrganizationWriteAccess *bool                                          `json:"blockCrossOrganizationWriteAccess,omitempty"`
	EncryptionMethod                  *BitLockerRemovableDrivePolicyEncryptionMethod `json:"encryptionMethod,omitempty"`
	ODataType                         *string                                        `json:"@odata.type,omitempty"`
	RequireEncryptionForWriteAccess   *bool                                          `json:"requireEncryptionForWriteAccess,omitempty"`
}
