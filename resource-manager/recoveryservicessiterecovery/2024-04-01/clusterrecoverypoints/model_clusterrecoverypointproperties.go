package clusterrecoverypoints

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClusterRecoveryPointProperties struct {
	ProviderSpecificDetails ClusterProviderSpecificRecoveryPointDetails `json:"providerSpecificDetails"`
	RecoveryPointTime       *string                                     `json:"recoveryPointTime,omitempty"`
	RecoveryPointType       *ClusterRecoveryPointType                   `json:"recoveryPointType,omitempty"`
}

var _ json.Unmarshaler = &ClusterRecoveryPointProperties{}

func (s *ClusterRecoveryPointProperties) UnmarshalJSON(bytes []byte) error {
	type alias ClusterRecoveryPointProperties
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into ClusterRecoveryPointProperties: %+v", err)
	}

	s.RecoveryPointTime = decoded.RecoveryPointTime
	s.RecoveryPointType = decoded.RecoveryPointType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ClusterRecoveryPointProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["providerSpecificDetails"]; ok {
		impl, err := unmarshalClusterProviderSpecificRecoveryPointDetailsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ProviderSpecificDetails' for 'ClusterRecoveryPointProperties': %+v", err)
		}
		s.ProviderSpecificDetails = impl
	}
	return nil
}
