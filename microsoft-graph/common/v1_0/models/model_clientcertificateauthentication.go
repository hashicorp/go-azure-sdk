package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClientCertificateAuthentication struct {
	CertificateList *[]Pkcs12CertificateInformation `json:"certificateList,omitempty"`
	ODataType       *string                         `json:"@odata.type,omitempty"`
}
