package caches

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SmbSettings struct {
	SmbAccessBasedEnumeration *SmbAccessBasedEnumeration `json:"smbAccessBasedEnumeration,omitempty"`
	SmbEncryption             *SmbEncryptionState        `json:"smbEncryption,omitempty"`
	SmbNonBrowsable           *SmbNonBrowsable           `json:"smbNonBrowsable,omitempty"`
}
