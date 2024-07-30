package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecuritySslCertificate struct {
	ExpirationDateTime *string                       `json:"expirationDateTime,omitempty"`
	Fingerprint        *string                       `json:"fingerprint,omitempty"`
	FirstSeenDateTime  *string                       `json:"firstSeenDateTime,omitempty"`
	Id                 *string                       `json:"id,omitempty"`
	IssueDateTime      *string                       `json:"issueDateTime,omitempty"`
	Issuer             *SecuritySslCertificateEntity `json:"issuer,omitempty"`
	LastSeenDateTime   *string                       `json:"lastSeenDateTime,omitempty"`
	ODataType          *string                       `json:"@odata.type,omitempty"`
	RelatedHosts       *[]SecurityHost               `json:"relatedHosts,omitempty"`
	SerialNumber       *string                       `json:"serialNumber,omitempty"`
	Sha1               *string                       `json:"sha1,omitempty"`
	Subject            *SecuritySslCertificateEntity `json:"subject,omitempty"`
}
