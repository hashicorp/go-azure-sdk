package lifecycleworkflowdeleteditemworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnlineMeeting struct {
	AllowAttendeeToEnableCamera    *bool                                        `json:"allowAttendeeToEnableCamera,omitempty"`
	AllowAttendeeToEnableMic       *bool                                        `json:"allowAttendeeToEnableMic,omitempty"`
	AllowMeetingChat               *OnlineMeetingAllowMeetingChat               `json:"allowMeetingChat,omitempty"`
	AllowParticipantsToChangeName  *bool                                        `json:"allowParticipantsToChangeName,omitempty"`
	AllowTeamworkReactions         *bool                                        `json:"allowTeamworkReactions,omitempty"`
	AllowedPresenters              *OnlineMeetingAllowedPresenters              `json:"allowedPresenters,omitempty"`
	AttendanceReports              *[]MeetingAttendanceReport                   `json:"attendanceReports,omitempty"`
	AttendeeReport                 *string                                      `json:"attendeeReport,omitempty"`
	AudioConferencing              *AudioConferencing                           `json:"audioConferencing,omitempty"`
	BroadcastSettings              *BroadcastMeetingSettings                    `json:"broadcastSettings,omitempty"`
	ChatInfo                       *ChatInfo                                    `json:"chatInfo,omitempty"`
	CreationDateTime               *string                                      `json:"creationDateTime,omitempty"`
	EndDateTime                    *string                                      `json:"endDateTime,omitempty"`
	ExternalId                     *string                                      `json:"externalId,omitempty"`
	Id                             *string                                      `json:"id,omitempty"`
	IsBroadcast                    *bool                                        `json:"isBroadcast,omitempty"`
	IsEntryExitAnnounced           *bool                                        `json:"isEntryExitAnnounced,omitempty"`
	JoinInformation                *ItemBody                                    `json:"joinInformation,omitempty"`
	JoinMeetingIdSettings          *JoinMeetingIdSettings                       `json:"joinMeetingIdSettings,omitempty"`
	JoinWebUrl                     *string                                      `json:"joinWebUrl,omitempty"`
	LobbyBypassSettings            *LobbyBypassSettings                         `json:"lobbyBypassSettings,omitempty"`
	ODataType                      *string                                      `json:"@odata.type,omitempty"`
	Participants                   *MeetingParticipants                         `json:"participants,omitempty"`
	RecordAutomatically            *bool                                        `json:"recordAutomatically,omitempty"`
	Recordings                     *[]CallRecording                             `json:"recordings,omitempty"`
	ShareMeetingChatHistoryDefault *OnlineMeetingShareMeetingChatHistoryDefault `json:"shareMeetingChatHistoryDefault,omitempty"`
	StartDateTime                  *string                                      `json:"startDateTime,omitempty"`
	Subject                        *string                                      `json:"subject,omitempty"`
	Transcripts                    *[]CallTranscript                            `json:"transcripts,omitempty"`
	VideoTeleconferenceId          *string                                      `json:"videoTeleconferenceId,omitempty"`
	WatermarkProtection            *WatermarkProtectionValues                   `json:"watermarkProtection,omitempty"`
}
