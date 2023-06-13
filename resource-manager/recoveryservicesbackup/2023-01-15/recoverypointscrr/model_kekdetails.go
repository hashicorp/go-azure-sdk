package recoverypointscrr

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KEKDetails struct {
	KeyBackupData *string `json:"keyBackupData,omitempty"`
	KeyUrl        *string `json:"keyUrl,omitempty"`
	KeyVaultId    *string `json:"keyVaultId,omitempty"`
}
