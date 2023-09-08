package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnterpriseCodeSigningCertificate struct {
	Content            *string                                 `json:"content,omitempty"`
	ExpirationDateTime *string                                 `json:"expirationDateTime,omitempty"`
	Id                 *string                                 `json:"id,omitempty"`
	Issuer             *string                                 `json:"issuer,omitempty"`
	IssuerName         *string                                 `json:"issuerName,omitempty"`
	ODataType          *string                                 `json:"@odata.type,omitempty"`
	Status             *EnterpriseCodeSigningCertificateStatus `json:"status,omitempty"`
	Subject            *string                                 `json:"subject,omitempty"`
	SubjectName        *string                                 `json:"subjectName,omitempty"`
	UploadDateTime     *string                                 `json:"uploadDateTime,omitempty"`
}
