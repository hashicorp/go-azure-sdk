package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DayNote struct {
	CreatedBy            *IdentitySet `json:"createdBy,omitempty"`
	CreatedDateTime      *string      `json:"createdDateTime,omitempty"`
	DayNoteDate          *string      `json:"dayNoteDate,omitempty"`
	DraftDayNote         *ItemBody    `json:"draftDayNote,omitempty"`
	Id                   *string      `json:"id,omitempty"`
	LastModifiedBy       *IdentitySet `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime *string      `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string      `json:"@odata.type,omitempty"`
	SharedDayNote        *ItemBody    `json:"sharedDayNote,omitempty"`
}
