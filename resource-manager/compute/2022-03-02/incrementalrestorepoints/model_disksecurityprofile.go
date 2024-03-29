package incrementalrestorepoints

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiskSecurityProfile struct {
	SecureVMDiskEncryptionSetId *string            `json:"secureVMDiskEncryptionSetId,omitempty"`
	SecurityType                *DiskSecurityTypes `json:"securityType,omitempty"`
}
