package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttackSimulationUser struct {
	DisplayName     *string `json:"displayName,omitempty"`
	Email           *string `json:"email,omitempty"`
	ODataType       *string `json:"@odata.type,omitempty"`
	OutOfOfficeDays *int64  `json:"outOfOfficeDays,omitempty"`
	UserId          *string `json:"userId,omitempty"`
}
