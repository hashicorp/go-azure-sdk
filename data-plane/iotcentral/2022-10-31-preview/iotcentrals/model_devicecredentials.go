package iotcentrals

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceCredentials struct {
	IdScope      string        `json:"idScope"`
	SymmetricKey *SymmetricKey `json:"symmetricKey,omitempty"`
	Tpm          *Tpm          `json:"tpm,omitempty"`
	X509         *X509         `json:"x509,omitempty"`
}
