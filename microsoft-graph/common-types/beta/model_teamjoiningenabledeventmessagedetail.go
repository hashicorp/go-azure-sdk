package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EventMessageDetail = TeamJoiningEnabledEventMessageDetail{}

type TeamJoiningEnabledEventMessageDetail struct {
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

func (s TeamJoiningEnabledEventMessageDetail) EventMessageDetail() BaseEventMessageDetailImpl {
	return BaseEventMessageDetailImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = TeamJoiningEnabledEventMessageDetail{}

func (s TeamJoiningEnabledEventMessageDetail) MarshalJSON() ([]byte, error) {
	type wrapper TeamJoiningEnabledEventMessageDetail
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TeamJoiningEnabledEventMessageDetail: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TeamJoiningEnabledEventMessageDetail: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.teamJoiningEnabledEventMessageDetail"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TeamJoiningEnabledEventMessageDetail: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &TeamJoiningEnabledEventMessageDetail{}

func (s *TeamJoiningEnabledEventMessageDetail) UnmarshalJSON(bytes []byte) error {
	type alias TeamJoiningEnabledEventMessageDetail
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into TeamJoiningEnabledEventMessageDetail: %+v", err)
	}

	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.TeamId = decoded.TeamId

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling TeamJoiningEnabledEventMessageDetail into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["initiator"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Initiator' for 'TeamJoiningEnabledEventMessageDetail': %+v", err)
		}
		s.Initiator = impl
	}
	return nil
}
