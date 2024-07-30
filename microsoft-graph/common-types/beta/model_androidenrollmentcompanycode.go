package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidEnrollmentCompanyCode struct {
	EnrollmentToken *string      `json:"enrollmentToken,omitempty"`
	ODataType       *string      `json:"@odata.type,omitempty"`
	QrCodeContent   *string      `json:"qrCodeContent,omitempty"`
	QrCodeImage     *MimeContent `json:"qrCodeImage,omitempty"`
}
