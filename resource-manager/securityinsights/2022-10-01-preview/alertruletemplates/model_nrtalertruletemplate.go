package alertruletemplates

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AlertRuleTemplate = NrtAlertRuleTemplate{}

type NrtAlertRuleTemplate struct {
	Properties AlertRuleTemplate `json:"properties"`

	// Fields inherited from AlertRuleTemplate
	Id         *string                `json:"id,omitempty"`
	Name       *string                `json:"name,omitempty"`
	SystemData *systemdata.SystemData `json:"systemData,omitempty"`
	Type       *string                `json:"type,omitempty"`
}

var _ json.Marshaler = NrtAlertRuleTemplate{}

func (s NrtAlertRuleTemplate) MarshalJSON() ([]byte, error) {
	type wrapper NrtAlertRuleTemplate
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling NrtAlertRuleTemplate: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling NrtAlertRuleTemplate: %+v", err)
	}
	decoded["kind"] = "NRT"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling NrtAlertRuleTemplate: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &NrtAlertRuleTemplate{}

func (s *NrtAlertRuleTemplate) UnmarshalJSON(bytes []byte) error {
	type alias NrtAlertRuleTemplate
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into NrtAlertRuleTemplate: %+v", err)
	}

	s.Id = decoded.Id
	s.Name = decoded.Name
	s.SystemData = decoded.SystemData
	s.Type = decoded.Type

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling NrtAlertRuleTemplate into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["properties"]; ok {
		impl, err := unmarshalAlertRuleTemplateImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Properties' for 'NrtAlertRuleTemplate': %+v", err)
		}
		s.Properties = impl
	}
	return nil
}
