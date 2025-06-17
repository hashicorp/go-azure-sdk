package cloudcertificationauthority

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RevokeCloudCertificationAuthorityCertificateRequest struct {
	CertificationAuthorityVersion *int64 `json:"certificationAuthorityVersion,omitempty"`
}
