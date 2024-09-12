package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EventMessageDetail = MeetingPolicyUpdatedEventMessageDetail{}

type MeetingPolicyUpdatedEventMessageDetail struct {
	// Initiator of the event.
	Initiator IdentitySet `json:"initiator"`

	// Represents whether the meeting chat is enabled or not.
	MeetingChatEnabled nullable.Type[bool] `json:"meetingChatEnabled,omitempty"`

	// Unique identifier of the meeting chat.
	MeetingChatId nullable.Type[string] `json:"meetingChatId,omitempty"`

	// Fields inherited from EventMessageDetail

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s MeetingPolicyUpdatedEventMessageDetail) EventMessageDetail() BaseEventMessageDetailImpl {
	return BaseEventMessageDetailImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MeetingPolicyUpdatedEventMessageDetail{}

func (s MeetingPolicyUpdatedEventMessageDetail) MarshalJSON() ([]byte, error) {
	type wrapper MeetingPolicyUpdatedEventMessageDetail
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MeetingPolicyUpdatedEventMessageDetail: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MeetingPolicyUpdatedEventMessageDetail: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.meetingPolicyUpdatedEventMessageDetail"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MeetingPolicyUpdatedEventMessageDetail: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &MeetingPolicyUpdatedEventMessageDetail{}

func (s *MeetingPolicyUpdatedEventMessageDetail) UnmarshalJSON(bytes []byte) error {
	type alias MeetingPolicyUpdatedEventMessageDetail
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into MeetingPolicyUpdatedEventMessageDetail: %+v", err)
	}

	s.MeetingChatEnabled = decoded.MeetingChatEnabled
	s.MeetingChatId = decoded.MeetingChatId
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling MeetingPolicyUpdatedEventMessageDetail into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["initiator"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Initiator' for 'MeetingPolicyUpdatedEventMessageDetail': %+v", err)
		}
		s.Initiator = impl
	}
	return nil
}
