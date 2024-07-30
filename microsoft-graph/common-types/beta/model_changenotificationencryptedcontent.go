package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChangeNotificationEncryptedContent struct {
	Data                            *string `json:"data,omitempty"`
	DataKey                         *string `json:"dataKey,omitempty"`
	DataSignature                   *string `json:"dataSignature,omitempty"`
	EncryptionCertificateId         *string `json:"encryptionCertificateId,omitempty"`
	EncryptionCertificateThumbprint *string `json:"encryptionCertificateThumbprint,omitempty"`
	ODataType                       *string `json:"@odata.type,omitempty"`
}
