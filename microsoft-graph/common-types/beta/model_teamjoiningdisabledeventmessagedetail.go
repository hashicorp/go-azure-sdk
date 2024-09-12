package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EventMessageDetail = TeamJoiningDisabledEventMessageDetail{}

type TeamJoiningDisabledEventMessageDetail struct {
	// Initiator of the event.
	Initiator IdentitySet `json:"initiator"`

	// Unique identifier of the team.
	TeamId nullable.Type[string] `json:"teamId,omitempty"`

	// Fields inherited from EventMessageDetail

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s TeamJoiningDisabledEventMessageDetail) EventMessageDetail() BaseEventMessageDetailImpl {
	return BaseEventMessageDetailImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = TeamJoiningDisabledEventMessageDetail{}

func (s TeamJoiningDisabledEventMessageDetail) MarshalJSON() ([]byte, error) {
	type wrapper TeamJoiningDisabledEventMessageDetail
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TeamJoiningDisabledEventMessageDetail: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TeamJoiningDisabledEventMessageDetail: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.teamJoiningDisabledEventMessageDetail"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TeamJoiningDisabledEventMessageDetail: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &TeamJoiningDisabledEventMessageDetail{}

func (s *TeamJoiningDisabledEventMessageDetail) UnmarshalJSON(bytes []byte) error {
	type alias TeamJoiningDisabledEventMessageDetail
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into TeamJoiningDisabledEventMessageDetail: %+v", err)
	}

	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.TeamId = decoded.TeamId

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling TeamJoiningDisabledEventMessageDetail into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["initiator"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Initiator' for 'TeamJoiningDisabledEventMessageDetail': %+v", err)
		}
		s.Initiator = impl
	}
	return nil
}
