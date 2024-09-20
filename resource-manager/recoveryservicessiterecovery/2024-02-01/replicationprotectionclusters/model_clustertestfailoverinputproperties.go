package replicationprotectionclusters

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClusterTestFailoverInputProperties struct {
	FailoverDirection       *FailoverDirection                       `json:"failoverDirection,omitempty"`
	NetworkId               *string                                  `json:"networkId,omitempty"`
	NetworkType             *string                                  `json:"networkType,omitempty"`
	ProviderSpecificDetails ClusterTestFailoverProviderSpecificInput `json:"providerSpecificDetails"`
}

var _ json.Unmarshaler = &ClusterTestFailoverInputProperties{}

func (s *ClusterTestFailoverInputProperties) UnmarshalJSON(bytes []byte) error {
	type alias ClusterTestFailoverInputProperties
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into ClusterTestFailoverInputProperties: %+v", err)
	}

	s.FailoverDirection = decoded.FailoverDirection
	s.NetworkId = decoded.NetworkId
	s.NetworkType = decoded.NetworkType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ClusterTestFailoverInputProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["providerSpecificDetails"]; ok {
		impl, err := UnmarshalClusterTestFailoverProviderSpecificInputImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ProviderSpecificDetails' for 'ClusterTestFailoverInputProperties': %+v", err)
		}
		s.ProviderSpecificDetails = impl
	}
	return nil
}
