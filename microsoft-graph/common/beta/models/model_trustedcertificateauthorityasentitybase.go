package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TrustedCertificateAuthorityAsEntityBase struct {
	DeletedDateTime               *string                         `json:"deletedDateTime,omitempty"`
	Id                            *string                         `json:"id,omitempty"`
	ODataType                     *string                         `json:"@odata.type,omitempty"`
	TrustedCertificateAuthorities *[]CertificateAuthorityAsEntity `json:"trustedCertificateAuthorities,omitempty"`
}
