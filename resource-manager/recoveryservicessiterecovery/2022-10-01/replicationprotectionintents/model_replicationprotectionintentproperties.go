package replicationprotectionintents

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReplicationProtectionIntentProperties struct {
	CreationTimeUTC         *string                                              `json:"creationTimeUTC,omitempty"`
	FriendlyName            *string                                              `json:"friendlyName,omitempty"`
	IsActive                *bool                                                `json:"isActive,omitempty"`
	JobId                   *string                                              `json:"jobId,omitempty"`
	JobState                *string                                              `json:"jobState,omitempty"`
	ProviderSpecificDetails *ReplicationProtectionIntentProviderSpecificSettings `json:"providerSpecificDetails,omitempty"`
}

var _ json.Unmarshaler = &ReplicationProtectionIntentProperties{}

func (s *ReplicationProtectionIntentProperties) UnmarshalJSON(bytes []byte) error {
	type alias ReplicationProtectionIntentProperties
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into ReplicationProtectionIntentProperties: %+v", err)
	}

	s.CreationTimeUTC = decoded.CreationTimeUTC
	s.FriendlyName = decoded.FriendlyName
	s.IsActive = decoded.IsActive
	s.JobId = decoded.JobId
	s.JobState = decoded.JobState

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ReplicationProtectionIntentProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["providerSpecificDetails"]; ok {
		impl, err := unmarshalReplicationProtectionIntentProviderSpecificSettingsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ProviderSpecificDetails' for 'ReplicationProtectionIntentProperties': %+v", err)
		}
		s.ProviderSpecificDetails = &impl
	}
	return nil
}
