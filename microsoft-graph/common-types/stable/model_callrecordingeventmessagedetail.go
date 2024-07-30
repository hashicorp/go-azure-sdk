package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordingEventMessageDetail struct {
	CallId                   *string                                             `json:"callId,omitempty"`
	CallRecordingDisplayName *string                                             `json:"callRecordingDisplayName,omitempty"`
	CallRecordingDuration    *string                                             `json:"callRecordingDuration,omitempty"`
	CallRecordingStatus      *CallRecordingEventMessageDetailCallRecordingStatus `json:"callRecordingStatus,omitempty"`
	CallRecordingUrl         *string                                             `json:"callRecordingUrl,omitempty"`
	Initiator                *IdentitySet                                        `json:"initiator,omitempty"`
	MeetingOrganizer         *IdentitySet                                        `json:"meetingOrganizer,omitempty"`
	ODataType                *string                                             `json:"@odata.type,omitempty"`
}
