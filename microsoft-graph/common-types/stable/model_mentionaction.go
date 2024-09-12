package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MentionAction struct {
	// The identities of the users mentioned in this action.
	Mentionees *[]IdentitySet `json:"mentionees,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

var _ json.Unmarshaler = &MentionAction{}

func (s *MentionAction) UnmarshalJSON(bytes []byte) error {
	type alias MentionAction
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into MentionAction: %+v", err)
	}

	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling MentionAction into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["mentionees"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Mentionees into list []json.RawMessage: %+v", err)
		}

		output := make([]IdentitySet, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalIdentitySetImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Mentionees' for 'MentionAction': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Mentionees = &output
	}
	return nil
}
