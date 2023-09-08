package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type X509CertificateUserBinding struct {
	ODataType            *string `json:"@odata.type,omitempty"`
	Priority             *int64  `json:"priority,omitempty"`
	UserProperty         *string `json:"userProperty,omitempty"`
	X509CertificateField *string `json:"x509CertificateField,omitempty"`
}
