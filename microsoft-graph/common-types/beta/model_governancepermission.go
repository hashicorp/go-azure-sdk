package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GovernancePermission struct {
	AccessLevel *string `json:"accessLevel,omitempty"`
	IsActive    *bool   `json:"isActive,omitempty"`
	IsEligible  *bool   `json:"isEligible,omitempty"`
	ODataType   *string `json:"@odata.type,omitempty"`
}
