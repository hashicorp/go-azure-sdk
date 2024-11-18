package replicationprotectionclusters

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplyClusterRecoveryPointInputProperties struct {
	ClusterRecoveryPointId       *string                                        `json:"clusterRecoveryPointId,omitempty"`
	IndividualNodeRecoveryPoints *[]string                                      `json:"individualNodeRecoveryPoints,omitempty"`
	ProviderSpecificDetails      ApplyClusterRecoveryPointProviderSpecificInput `json:"providerSpecificDetails"`
}

var _ json.Unmarshaler = &ApplyClusterRecoveryPointInputProperties{}

func (s *ApplyClusterRecoveryPointInputProperties) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ClusterRecoveryPointId       *string   `json:"clusterRecoveryPointId,omitempty"`
		IndividualNodeRecoveryPoints *[]string `json:"individualNodeRecoveryPoints,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ClusterRecoveryPointId = decoded.ClusterRecoveryPointId
	s.IndividualNodeRecoveryPoints = decoded.IndividualNodeRecoveryPoints

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ApplyClusterRecoveryPointInputProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["providerSpecificDetails"]; ok {
		impl, err := UnmarshalApplyClusterRecoveryPointProviderSpecificInputImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ProviderSpecificDetails' for 'ApplyClusterRecoveryPointInputProperties': %+v", err)
		}
		s.ProviderSpecificDetails = impl
	}

	return nil
}
