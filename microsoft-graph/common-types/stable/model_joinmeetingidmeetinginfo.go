package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MeetingInfo = JoinMeetingIdMeetingInfo{}

type JoinMeetingIdMeetingInfo struct {
	// The ID used to join the meeting.
	JoinMeetingId *string `json:"joinMeetingId,omitempty"`

	// The passcode used to join the meeting. Optional.
	Passcode nullable.Type[string] `json:"passcode,omitempty"`

	// Fields inherited from MeetingInfo

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s JoinMeetingIdMeetingInfo) MeetingInfo() BaseMeetingInfoImpl {
	return BaseMeetingInfoImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = JoinMeetingIdMeetingInfo{}

func (s JoinMeetingIdMeetingInfo) MarshalJSON() ([]byte, error) {
	type wrapper JoinMeetingIdMeetingInfo
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling JoinMeetingIdMeetingInfo: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling JoinMeetingIdMeetingInfo: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.joinMeetingIdMeetingInfo"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling JoinMeetingIdMeetingInfo: %+v", err)
	}

	return encoded, nil
}
