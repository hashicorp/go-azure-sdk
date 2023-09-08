package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Participant struct {
	Id                   *string                  `json:"id,omitempty"`
	Info                 *ParticipantInfo         `json:"info,omitempty"`
	IsInLobby            *bool                    `json:"isInLobby,omitempty"`
	IsMuted              *bool                    `json:"isMuted,omitempty"`
	MediaStreams         *[]MediaStream           `json:"mediaStreams,omitempty"`
	Metadata             *string                  `json:"metadata,omitempty"`
	ODataType            *string                  `json:"@odata.type,omitempty"`
	RecordingInfo        *RecordingInfo           `json:"recordingInfo,omitempty"`
	RestrictedExperience *OnlineMeetingRestricted `json:"restrictedExperience,omitempty"`
}
