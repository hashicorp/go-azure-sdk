package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeetingParticipantInfo struct {
	Identity  *IdentitySet                `json:"identity,omitempty"`
	ODataType *string                     `json:"@odata.type,omitempty"`
	Role      *MeetingParticipantInfoRole `json:"role,omitempty"`
	Upn       *string                     `json:"upn,omitempty"`
}
