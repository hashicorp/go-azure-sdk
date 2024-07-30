package virtualendpointprovisioningpolicy

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
	AnonymizeIdentityForRoles      *VirtualEventSessionAnonymizeIdentityForRoles      `json:"anonymizeIdentityForRoles,omitempty"`
	AttendanceReports              *[]MeetingAttendanceReport                         `json:"attendanceReports,omitempty"`
	AudioConferencing              *AudioConferencing                                 `json:"audioConferencing,omitempty"`
	ChatInfo                       *ChatInfo                                          `json:"chatInfo,omitempty"`
	ChatRestrictions               *ChatRestrictions                                  `json:"chatRestrictions,omitempty"`
	EndDateTime                    *DateTimeTimeZone                                  `json:"endDateTime,omitempty"`
	Id                             *string                                            `json:"id,omitempty"`
	IsEndToEndEncryptionEnabled    *bool                                              `json:"isEndToEndEncryptionEnabled,omitempty"`
	IsEntryExitAnnounced           *bool                                              `json:"isEntryExitAnnounced,omitempty"`
	JoinInformation                *ItemBody                                          `json:"joinInformation,omitempty"`
	JoinMeetingIdSettings          *JoinMeetingIdSettings                             `json:"joinMeetingIdSettings,omitempty"`
	JoinWebUrl                     *string                                            `json:"joinWebUrl,omitempty"`
	LobbyBypassSettings            *LobbyBypassSettings                               `json:"lobbyBypassSettings,omitempty"`
	ODataType                      *string                                            `json:"@odata.type,omitempty"`
	Presenters                     *[]VirtualEventPresenter                           `json:"presenters,omitempty"`
	RecordAutomatically            *bool                                              `json:"recordAutomatically,omitempty"`
	Registrations                  *[]VirtualEventRegistration                        `json:"registrations,omitempty"`
	ShareMeetingChatHistoryDefault *VirtualEventSessionShareMeetingChatHistoryDefault `json:"shareMeetingChatHistoryDefault,omitempty"`
	StartDateTime                  *DateTimeTimeZone                                  `json:"startDateTime,omitempty"`
	Subject                        *string                                            `json:"subject,omitempty"`
	VideoTeleconferenceId          *string                                            `json:"videoTeleconferenceId,omitempty"`
	WatermarkProtection            *WatermarkProtectionValues                         `json:"watermarkProtection,omitempty"`
}
