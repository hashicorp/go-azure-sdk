package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ParticipantJoiningResponse interface {
	ParticipantJoiningResponse() BaseParticipantJoiningResponseImpl
}

var _ ParticipantJoiningResponse = BaseParticipantJoiningResponseImpl{}

type BaseParticipantJoiningResponseImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s BaseParticipantJoiningResponseImpl) ParticipantJoiningResponse() BaseParticipantJoiningResponseImpl {
	return s
}

var _ ParticipantJoiningResponse = RawParticipantJoiningResponseImpl{}

// RawParticipantJoiningResponseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawParticipantJoiningResponseImpl struct {
	participantJoiningResponse BaseParticipantJoiningResponseImpl
	Type                       string
	Values                     map[string]interface{}
}

func (s RawParticipantJoiningResponseImpl) ParticipantJoiningResponse() BaseParticipantJoiningResponseImpl {
	return s.participantJoiningResponse
}

func UnmarshalParticipantJoiningResponseImplementation(input []byte) (ParticipantJoiningResponse, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ParticipantJoiningResponse into map[string]interface: %+v", err)
	}

	value, ok := temp["@odata.type"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.acceptJoinResponse") {
		var out AcceptJoinResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AcceptJoinResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.inviteNewBotResponse") {
		var out InviteNewBotResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into InviteNewBotResponse: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.rejectJoinResponse") {
		var out RejectJoinResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RejectJoinResponse: %+v", err)
		}
		return out, nil
	}

	var parent BaseParticipantJoiningResponseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseParticipantJoiningResponseImpl: %+v", err)
	}

	return RawParticipantJoiningResponseImpl{
		participantJoiningResponse: parent,
		Type:                       value,
		Values:                     temp,
	}, nil

}
