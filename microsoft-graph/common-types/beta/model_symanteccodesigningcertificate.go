package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SymantecCodeSigningCertificate struct {
	Content            *string                               `json:"content,omitempty"`
	ExpirationDateTime *string                               `json:"expirationDateTime,omitempty"`
	Id                 *string                               `json:"id,omitempty"`
	Issuer             *string                               `json:"issuer,omitempty"`
	IssuerName         *string                               `json:"issuerName,omitempty"`
	ODataType          *string                               `json:"@odata.type,omitempty"`
	Password           *string                               `json:"password,omitempty"`
	Status             *SymantecCodeSigningCertificateStatus `json:"status,omitempty"`
	Subject            *string                               `json:"subject,omitempty"`
	SubjectName        *string                               `json:"subjectName,omitempty"`
	UploadDateTime     *string                               `json:"uploadDateTime,omitempty"`
}
