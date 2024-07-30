package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityHostSslCertificate struct {
	FirstSeenDateTime *string                           `json:"firstSeenDateTime,omitempty"`
	Host              *SecurityHost                     `json:"host,omitempty"`
	Id                *string                           `json:"id,omitempty"`
	LastSeenDateTime  *string                           `json:"lastSeenDateTime,omitempty"`
	ODataType         *string                           `json:"@odata.type,omitempty"`
	Ports             *[]SecurityHostSslCertificatePort `json:"ports,omitempty"`
	SslCertificate    *SecuritySslCertificate           `json:"sslCertificate,omitempty"`
}
