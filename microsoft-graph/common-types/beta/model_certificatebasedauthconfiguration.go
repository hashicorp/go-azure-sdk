package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CertificateBasedAuthConfiguration struct {
	CertificateAuthorities *[]CertificateAuthority `json:"certificateAuthorities,omitempty"`
	Id                     *string                 `json:"id,omitempty"`
	ODataType              *string                 `json:"@odata.type,omitempty"`
}
