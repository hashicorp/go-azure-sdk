package serviceprincipal

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddServicePrincipalTokenSigningCertificateRequest struct {
	DisplayName *string `json:"displayName,omitempty"`
	EndDateTime *string `json:"endDateTime,omitempty"`
}
