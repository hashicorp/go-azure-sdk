package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeetingCapability struct {
	AllowAnonymousUsersToDialOut      *bool                               `json:"allowAnonymousUsersToDialOut,omitempty"`
	AllowAnonymousUsersToStartMeeting *bool                               `json:"allowAnonymousUsersToStartMeeting,omitempty"`
	AutoAdmittedUsers                 *MeetingCapabilityAutoAdmittedUsers `json:"autoAdmittedUsers,omitempty"`
	ODataType                         *string                             `json:"@odata.type,omitempty"`
}
