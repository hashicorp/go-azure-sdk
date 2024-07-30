package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamRenamedEventMessageDetail struct {
	Initiator       *IdentitySet `json:"initiator,omitempty"`
	ODataType       *string      `json:"@odata.type,omitempty"`
	TeamDisplayName *string      `json:"teamDisplayName,omitempty"`
	TeamId          *string      `json:"teamId,omitempty"`
}
