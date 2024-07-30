package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessEnumeratedExternalTenants struct {
	Members        *[]string                                                 `json:"members,omitempty"`
	MembershipKind *ConditionalAccessEnumeratedExternalTenantsMembershipKind `json:"membershipKind,omitempty"`
	ODataType      *string                                                   `json:"@odata.type,omitempty"`
}
