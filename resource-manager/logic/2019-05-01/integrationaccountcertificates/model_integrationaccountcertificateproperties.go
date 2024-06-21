package integrationaccountcertificates

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntegrationAccountCertificateProperties struct {
	ChangedTime       *string               `json:"changedTime,omitempty"`
	CreatedTime       *string               `json:"createdTime,omitempty"`
	Key               *KeyVaultKeyReference `json:"key,omitempty"`
	Metadata          *interface{}          `json:"metadata,omitempty"`
	PublicCertificate *string               `json:"publicCertificate,omitempty"`
}
