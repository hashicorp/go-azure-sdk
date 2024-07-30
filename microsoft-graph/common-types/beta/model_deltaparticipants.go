package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeltaParticipants struct {
	Id             *string        `json:"id,omitempty"`
	ODataType      *string        `json:"@odata.type,omitempty"`
	Participants   *[]Participant `json:"participants,omitempty"`
	SequenceNumber *int64         `json:"sequenceNumber,omitempty"`
}
