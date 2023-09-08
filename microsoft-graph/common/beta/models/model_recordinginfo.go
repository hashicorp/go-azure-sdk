package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecordingInfo struct {
	InitiatedBy     *ParticipantInfo              `json:"initiatedBy,omitempty"`
	Initiator       *IdentitySet                  `json:"initiator,omitempty"`
	ODataType       *string                       `json:"@odata.type,omitempty"`
	RecordingStatus *RecordingInfoRecordingStatus `json:"recordingStatus,omitempty"`
}
