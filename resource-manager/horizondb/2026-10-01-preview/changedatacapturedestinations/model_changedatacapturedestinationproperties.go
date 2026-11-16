package changedatacapturedestinations

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChangeDataCaptureDestinationProperties struct {
	ActiveIdentityAssignmentResourceId *string                                    `json:"activeIdentityAssignmentResourceId,omitempty"`
	ConfigurationCount                 int64                                      `json:"configurationCount"`
	DesiredState                       ChangeDataCaptureDesiredState              `json:"desiredState"`
	Destination                        ChangeDataCaptureDestinationDetails        `json:"destination"`
	EffectiveState                     ChangeDataCaptureDestinationEffectiveState `json:"effectiveState"`
	ProvisioningState                  *ProvisioningState                         `json:"provisioningState,omitempty"`
	Status                             ChangeDataCaptureStatus                    `json:"status"`
}

var _ json.Unmarshaler = &ChangeDataCaptureDestinationProperties{}

func (s *ChangeDataCaptureDestinationProperties) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ActiveIdentityAssignmentResourceId *string                                    `json:"activeIdentityAssignmentResourceId,omitempty"`
		ConfigurationCount                 int64                                      `json:"configurationCount"`
		DesiredState                       ChangeDataCaptureDesiredState              `json:"desiredState"`
		EffectiveState                     ChangeDataCaptureDestinationEffectiveState `json:"effectiveState"`
		ProvisioningState                  *ProvisioningState                         `json:"provisioningState,omitempty"`
		Status                             ChangeDataCaptureStatus                    `json:"status"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ActiveIdentityAssignmentResourceId = decoded.ActiveIdentityAssignmentResourceId
	s.ConfigurationCount = decoded.ConfigurationCount
	s.DesiredState = decoded.DesiredState
	s.EffectiveState = decoded.EffectiveState
	s.ProvisioningState = decoded.ProvisioningState
	s.Status = decoded.Status

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ChangeDataCaptureDestinationProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["destination"]; ok {
		impl, err := UnmarshalChangeDataCaptureDestinationDetailsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Destination' for 'ChangeDataCaptureDestinationProperties': %+v", err)
		}
		s.Destination = impl
	}

	return nil
}
