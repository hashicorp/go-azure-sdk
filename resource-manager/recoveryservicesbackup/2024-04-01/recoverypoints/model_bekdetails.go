package recoverypoints

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BEKDetails struct {
	SecretData    *string `json:"secretData,omitempty"`
	SecretUrl     *string `json:"secretUrl,omitempty"`
	SecretVaultId *string `json:"secretVaultId,omitempty"`
}
