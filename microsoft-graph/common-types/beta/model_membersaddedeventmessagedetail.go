package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MembersAddedEventMessageDetail struct {
	Initiator                   *IdentitySet            `json:"initiator,omitempty"`
	Members                     *[]TeamworkUserIdentity `json:"members,omitempty"`
	ODataType                   *string                 `json:"@odata.type,omitempty"`
	VisibleHistoryStartDateTime *string                 `json:"visibleHistoryStartDateTime,omitempty"`
}
