package pool

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VMDiskSecurityProfile struct {
	SecurityEncryptionType *SecurityEncryptionTypes `json:"securityEncryptionType,omitempty"`
}
