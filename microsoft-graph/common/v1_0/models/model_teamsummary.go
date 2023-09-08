package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamSummary struct {
	GuestsCount  *int64  `json:"guestsCount,omitempty"`
	MembersCount *int64  `json:"membersCount,omitempty"`
	ODataType    *string `json:"@odata.type,omitempty"`
	OwnersCount  *int64  `json:"ownersCount,omitempty"`
}
