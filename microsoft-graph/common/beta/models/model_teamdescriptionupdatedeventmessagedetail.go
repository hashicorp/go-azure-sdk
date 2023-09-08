package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamDescriptionUpdatedEventMessageDetail struct {
	Initiator       *IdentitySet `json:"initiator,omitempty"`
	ODataType       *string      `json:"@odata.type,omitempty"`
	TeamDescription *string      `json:"teamDescription,omitempty"`
	TeamId          *string      `json:"teamId,omitempty"`
}
