package alertruletemplates

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AlertRuleTemplate = MicrosoftSecurityIncidentCreationAlertRuleTemplate{}

type MicrosoftSecurityIncidentCreationAlertRuleTemplate struct {
	Properties *AlertRuleTemplate `json:"properties,omitempty"`

	// Fields inherited from AlertRuleTemplate
	Id         *string                `json:"id,omitempty"`
	Name       *string                `json:"name,omitempty"`
	SystemData *systemdata.SystemData `json:"systemData,omitempty"`
	Type       *string                `json:"type,omitempty"`
}

var _ json.Marshaler = MicrosoftSecurityIncidentCreationAlertRuleTemplate{}

func (s MicrosoftSecurityIncidentCreationAlertRuleTemplate) MarshalJSON() ([]byte, error) {
	type wrapper MicrosoftSecurityIncidentCreationAlertRuleTemplate
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MicrosoftSecurityIncidentCreationAlertRuleTemplate: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MicrosoftSecurityIncidentCreationAlertRuleTemplate: %+v", err)
	}
	decoded["kind"] = "MicrosoftSecurityIncidentCreation"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MicrosoftSecurityIncidentCreationAlertRuleTemplate: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &MicrosoftSecurityIncidentCreationAlertRuleTemplate{}

func (s *MicrosoftSecurityIncidentCreationAlertRuleTemplate) UnmarshalJSON(bytes []byte) error {
	type alias MicrosoftSecurityIncidentCreationAlertRuleTemplate
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into MicrosoftSecurityIncidentCreationAlertRuleTemplate: %+v", err)
	}

	s.Id = decoded.Id
	s.Name = decoded.Name
	s.SystemData = decoded.SystemData
	s.Type = decoded.Type

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling MicrosoftSecurityIncidentCreationAlertRuleTemplate into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["properties"]; ok {
		impl, err := unmarshalAlertRuleTemplateImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Properties' for 'MicrosoftSecurityIncidentCreationAlertRuleTemplate': %+v", err)
		}
		s.Properties = impl
	}
	return nil
}
