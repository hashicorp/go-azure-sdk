package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsInformationProtectionDataRecoveryCertificate struct {
	Certificate        *string `json:"certificate,omitempty"`
	Description        *string `json:"description,omitempty"`
	ExpirationDateTime *string `json:"expirationDateTime,omitempty"`
	ODataType          *string `json:"@odata.type,omitempty"`
	SubjectName        *string `json:"subjectName,omitempty"`
}
