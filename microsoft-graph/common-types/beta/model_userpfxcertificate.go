package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserPFXCertificate struct {
	CreatedDateTime      *string                            `json:"createdDateTime,omitempty"`
	EncryptedPfxBlob     *string                            `json:"encryptedPfxBlob,omitempty"`
	EncryptedPfxPassword *string                            `json:"encryptedPfxPassword,omitempty"`
	ExpirationDateTime   *string                            `json:"expirationDateTime,omitempty"`
	Id                   *string                            `json:"id,omitempty"`
	IntendedPurpose      *UserPFXCertificateIntendedPurpose `json:"intendedPurpose,omitempty"`
	KeyName              *string                            `json:"keyName,omitempty"`
	LastModifiedDateTime *string                            `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string                            `json:"@odata.type,omitempty"`
	PaddingScheme        *UserPFXCertificatePaddingScheme   `json:"paddingScheme,omitempty"`
	ProviderName         *string                            `json:"providerName,omitempty"`
	StartDateTime        *string                            `json:"startDateTime,omitempty"`
	Thumbprint           *string                            `json:"thumbprint,omitempty"`
	UserPrincipalName    *string                            `json:"userPrincipalName,omitempty"`
}
