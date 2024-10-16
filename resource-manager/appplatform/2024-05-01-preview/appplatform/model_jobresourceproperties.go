package appplatform

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobResourceProperties struct {
	ManagedComponentReferences *[]ManagedComponentReference  `json:"managedComponentReferences,omitempty"`
	ProvisioningState          *JobResourceProvisioningState `json:"provisioningState,omitempty"`
	Source                     UserSourceInfo                `json:"source"`
	Template                   *JobExecutionTemplate         `json:"template,omitempty"`
	TriggerConfig              JobTriggerConfig              `json:"triggerConfig"`
}

var _ json.Unmarshaler = &JobResourceProperties{}

func (s *JobResourceProperties) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ManagedComponentReferences *[]ManagedComponentReference  `json:"managedComponentReferences,omitempty"`
		ProvisioningState          *JobResourceProvisioningState `json:"provisioningState,omitempty"`
		Template                   *JobExecutionTemplate         `json:"template,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ManagedComponentReferences = decoded.ManagedComponentReferences
	s.ProvisioningState = decoded.ProvisioningState
	s.Template = decoded.Template

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling JobResourceProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["source"]; ok {
		impl, err := UnmarshalUserSourceInfoImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Source' for 'JobResourceProperties': %+v", err)
		}
		s.Source = impl
	}

	if v, ok := temp["triggerConfig"]; ok {
		impl, err := UnmarshalJobTriggerConfigImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'TriggerConfig' for 'JobResourceProperties': %+v", err)
		}
		s.TriggerConfig = impl
	}

	return nil
}
