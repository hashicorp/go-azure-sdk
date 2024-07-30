package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Call struct {
	ActiveModalities        *CallActiveModalities        `json:"activeModalities,omitempty"`
	AnsweredBy              *ParticipantInfo             `json:"answeredBy,omitempty"`
	AudioRoutingGroups      *[]AudioRoutingGroup         `json:"audioRoutingGroups,omitempty"`
	CallChainId             *string                      `json:"callChainId,omitempty"`
	CallOptions             *CallOptions                 `json:"callOptions,omitempty"`
	CallRoutes              *[]CallRoute                 `json:"callRoutes,omitempty"`
	CallbackUri             *string                      `json:"callbackUri,omitempty"`
	ChatInfo                *ChatInfo                    `json:"chatInfo,omitempty"`
	ContentSharingSessions  *[]ContentSharingSession     `json:"contentSharingSessions,omitempty"`
	Direction               *CallDirection               `json:"direction,omitempty"`
	Id                      *string                      `json:"id,omitempty"`
	IncomingContext         *IncomingContext             `json:"incomingContext,omitempty"`
	MediaConfig             *MediaConfig                 `json:"mediaConfig,omitempty"`
	MediaState              *CallMediaState              `json:"mediaState,omitempty"`
	MeetingCapability       *MeetingCapability           `json:"meetingCapability,omitempty"`
	MeetingInfo             *MeetingInfo                 `json:"meetingInfo,omitempty"`
	MyParticipantId         *string                      `json:"myParticipantId,omitempty"`
	ODataType               *string                      `json:"@odata.type,omitempty"`
	Operations              *[]CommsOperation            `json:"operations,omitempty"`
	Participants            *[]Participant               `json:"participants,omitempty"`
	RequestedModalities     *CallRequestedModalities     `json:"requestedModalities,omitempty"`
	ResultInfo              *ResultInfo                  `json:"resultInfo,omitempty"`
	RingingTimeoutInSeconds *int64                       `json:"ringingTimeoutInSeconds,omitempty"`
	RoutingPolicies         *CallRoutingPolicies         `json:"routingPolicies,omitempty"`
	Source                  *ParticipantInfo             `json:"source,omitempty"`
	State                   *CallState                   `json:"state,omitempty"`
	Subject                 *string                      `json:"subject,omitempty"`
	Targets                 *[]InvitationParticipantInfo `json:"targets,omitempty"`
	TenantId                *string                      `json:"tenantId,omitempty"`
	TerminationReason       *string                      `json:"terminationReason,omitempty"`
	ToneInfo                *ToneInfo                    `json:"toneInfo,omitempty"`
	Transcription           *CallTranscriptionInfo       `json:"transcription,omitempty"`
}
