package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesComplianceChange struct {
	CreatedDateTime *string                     `json:"createdDateTime,omitempty"`
	Id              *string                     `json:"id,omitempty"`
	IsRevoked       *bool                       `json:"isRevoked,omitempty"`
	ODataType       *string                     `json:"@odata.type,omitempty"`
	RevokedDateTime *string                     `json:"revokedDateTime,omitempty"`
	UpdatePolicy    *WindowsUpdatesUpdatePolicy `json:"updatePolicy,omitempty"`
}
