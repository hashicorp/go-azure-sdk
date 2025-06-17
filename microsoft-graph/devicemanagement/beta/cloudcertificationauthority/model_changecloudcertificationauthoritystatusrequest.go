package cloudcertificationauthority

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChangeCloudCertificationAuthorityStatusRequest struct {
	// Enum type of possible certification authority statuses. These statuses indicate whether a certification authority is
	// currently able to issue certificates or temporarily paused or permanently revoked.
	CertificationAuthorityStatus *beta.CloudCertificationAuthorityStatus `json:"certificationAuthorityStatus,omitempty"`

	CertificationAuthorityVersion *int64 `json:"certificationAuthorityVersion,omitempty"`
}
