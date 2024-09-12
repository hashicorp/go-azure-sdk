package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeetingInfo interface {
	MeetingInfo() BaseMeetingInfoImpl
}

var _ MeetingInfo = BaseMeetingInfoImpl{}

type BaseMeetingInfoImpl struct {
	AllowConversationWithoutHost nullable.Type[bool] `json:"allowConversationWithoutHost,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s BaseMeetingInfoImpl) MeetingInfo() BaseMeetingInfoImpl {
	return s
}

var _ MeetingInfo = RawMeetingInfoImpl{}

// RawMeetingInfoImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawMeetingInfoImpl struct {
	meetingInfo BaseMeetingInfoImpl
	Type        string
	Values      map[string]interface{}
}

func (s RawMeetingInfoImpl) MeetingInfo() BaseMeetingInfoImpl {
	return s.meetingInfo
}

func UnmarshalMeetingInfoImplementation(input []byte) (MeetingInfo, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling MeetingInfo into map[string]interface: %+v", err)
	}

	value, ok := temp["@odata.type"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.joinMeetingIdMeetingInfo") {
		var out JoinMeetingIdMeetingInfo
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into JoinMeetingIdMeetingInfo: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.organizerMeetingInfo") {
		var out OrganizerMeetingInfo
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OrganizerMeetingInfo: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.tokenMeetingInfo") {
		var out TokenMeetingInfo
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TokenMeetingInfo: %+v", err)
		}
		return out, nil
	}

	var parent BaseMeetingInfoImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseMeetingInfoImpl: %+v", err)
	}

	return RawMeetingInfoImpl{
		meetingInfo: parent,
		Type:        value,
		Values:      temp,
	}, nil

}
