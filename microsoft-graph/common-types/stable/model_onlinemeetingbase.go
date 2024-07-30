package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnlineMeetingBase struct {
	AllowAttendeeToEnableCamera    *bool                                            `json:"allowAttendeeToEnableCamera,omitempty"`
	AllowAttendeeToEnableMic       *bool                                            `json:"allowAttendeeToEnableMic,omitempty"`
	AllowMeetingChat               *OnlineMeetingBaseAllowMeetingChat               `json:"allowMeetingChat,omitempty"`
	AllowParticipantsToChangeName  *bool                                            `json:"allowParticipantsToChangeName,omitempty"`
	AllowTeamworkReactions         *bool                                            `json:"allowTeamworkReactions,omitempty"`
	AllowedPresenters              *OnlineMeetingBaseAllowedPresenters              `json:"allowedPresenters,omitempty"`
	AttendanceReports              *[]MeetingAttendanceReport                       `json:"attendanceReports,omitempty"`
	AudioConferencing              *AudioConferencing                               `json:"audioConferencing,omitempty"`
	ChatInfo                       *ChatInfo                                        `json:"chatInfo,omitempty"`
	Id                             *string                                          `json:"id,omitempty"`
	IsEntryExitAnnounced           *bool                                            `json:"isEntryExitAnnounced,omitempty"`
	JoinInformation                *ItemBody                                        `json:"joinInformation,omitempty"`
	JoinMeetingIdSettings          *JoinMeetingIdSettings                           `json:"joinMeetingIdSettings,omitempty"`
	JoinWebUrl                     *string                                          `json:"joinWebUrl,omitempty"`
	LobbyBypassSettings            *LobbyBypassSettings                             `json:"lobbyBypassSettings,omitempty"`
	ODataType                      *string                                          `json:"@odata.type,omitempty"`
	RecordAutomatically            *bool                                            `json:"recordAutomatically,omitempty"`
	ShareMeetingChatHistoryDefault *OnlineMeetingBaseShareMeetingChatHistoryDefault `json:"shareMeetingChatHistoryDefault,omitempty"`
	Subject                        *string                                          `json:"subject,omitempty"`
	VideoTeleconferenceId          *string                                          `json:"videoTeleconferenceId,omitempty"`
	WatermarkProtection            *WatermarkProtectionValues                       `json:"watermarkProtection,omitempty"`
}
