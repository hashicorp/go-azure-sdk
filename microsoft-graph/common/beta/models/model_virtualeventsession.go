package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEventSession struct {
	AllowAttendeeToEnableCamera    *bool                                              `json:"allowAttendeeToEnableCamera,omitempty"`
	AllowAttendeeToEnableMic       *bool                                              `json:"allowAttendeeToEnableMic,omitempty"`
	AllowMeetingChat               *VirtualEventSessionAllowMeetingChat               `json:"allowMeetingChat,omitempty"`
	AllowParticipantsToChangeName  *bool                                              `json:"allowParticipantsToChangeName,omitempty"`
	AllowRecording                 *bool                                              `json:"allowRecording,omitempty"`
	AllowTeamworkReactions         *bool                                              `json:"allowTeamworkReactions,omitempty"`
	AllowTranscription             *bool                                              `json:"allowTranscription,omitempty"`
	AllowedPresenters              *VirtualEventSessionAllowedPresenters              `json:"allowedPresenters,omitempty"`
	AlternativeRecording           *string                                            `json:"alternativeRecording,omitempty"`
	AnonymizeIdentityForRoles      *[]VirtualEventSessionAnonymizeIdentityForRoles    `json:"anonymizeIdentityForRoles,omitempty"`
	AttendanceReports              *[]MeetingAttendanceReport                         `json:"attendanceReports,omitempty"`
	AttendeeReport                 *string                                            `json:"attendeeReport,omitempty"`
	AudioConferencing              *AudioConferencing                                 `json:"audioConferencing,omitempty"`
	BroadcastRecording             *string                                            `json:"broadcastRecording,omitempty"`
	BroadcastSettings              *BroadcastMeetingSettings                          `json:"broadcastSettings,omitempty"`
	Capabilities                   *[]VirtualEventSessionCapabilities                 `json:"capabilities,omitempty"`
	ChatInfo                       *ChatInfo                                          `json:"chatInfo,omitempty"`
	ChatRestrictions               *ChatRestrictions                                  `json:"chatRestrictions,omitempty"`
	CreationDateTime               *string                                            `json:"creationDateTime,omitempty"`
	EndDateTime                    *string                                            `json:"endDateTime,omitempty"`
	ExternalId                     *string                                            `json:"externalId,omitempty"`
	Id                             *string                                            `json:"id,omitempty"`
	IsBroadcast                    *bool                                              `json:"isBroadcast,omitempty"`
	IsEntryExitAnnounced           *bool                                              `json:"isEntryExitAnnounced,omitempty"`
	JoinInformation                *ItemBody                                          `json:"joinInformation,omitempty"`
	JoinMeetingIdSettings          *JoinMeetingIdSettings                             `json:"joinMeetingIdSettings,omitempty"`
	JoinUrl                        *string                                            `json:"joinUrl,omitempty"`
	JoinWebUrl                     *string                                            `json:"joinWebUrl,omitempty"`
	LobbyBypassSettings            *LobbyBypassSettings                               `json:"lobbyBypassSettings,omitempty"`
	MeetingAttendanceReport        *MeetingAttendanceReport                           `json:"meetingAttendanceReport,omitempty"`
	ODataType                      *string                                            `json:"@odata.type,omitempty"`
	Participants                   *MeetingParticipants                               `json:"participants,omitempty"`
	RecordAutomatically            *bool                                              `json:"recordAutomatically,omitempty"`
	Recording                      *string                                            `json:"recording,omitempty"`
	Recordings                     *[]CallRecording                                   `json:"recordings,omitempty"`
	Registration                   *MeetingRegistration                               `json:"registration,omitempty"`
	Registrations                  *[]VirtualEventRegistration                        `json:"registrations,omitempty"`
	ShareMeetingChatHistoryDefault *VirtualEventSessionShareMeetingChatHistoryDefault `json:"shareMeetingChatHistoryDefault,omitempty"`
	StartDateTime                  *string                                            `json:"startDateTime,omitempty"`
	Subject                        *string                                            `json:"subject,omitempty"`
	Transcripts                    *[]CallTranscript                                  `json:"transcripts,omitempty"`
	VideoTeleconferenceId          *string                                            `json:"videoTeleconferenceId,omitempty"`
	WatermarkProtection            *WatermarkProtectionValues                         `json:"watermarkProtection,omitempty"`
}
