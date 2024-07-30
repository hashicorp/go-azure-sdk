package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamJoiningDisabledEventMessageDetail struct {
	Initiator *IdentitySet `json:"initiator,omitempty"`
	ODataType *string      `json:"@odata.type,omitempty"`
	TeamId    *string      `json:"teamId,omitempty"`
}
