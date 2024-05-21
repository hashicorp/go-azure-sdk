package replicationprotectionclusters

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClusterUnplannedFailoverInputProperties struct {
	FailoverDirection       *string                                       `json:"failoverDirection,omitempty"`
	ProviderSpecificDetails ClusterUnplannedFailoverProviderSpecificInput `json:"providerSpecificDetails"`
	SourceSiteOperations    *string                                       `json:"sourceSiteOperations,omitempty"`
}

var _ json.Unmarshaler = &ClusterUnplannedFailoverInputProperties{}

func (s *ClusterUnplannedFailoverInputProperties) UnmarshalJSON(bytes []byte) error {
	type alias ClusterUnplannedFailoverInputProperties
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into ClusterUnplannedFailoverInputProperties: %+v", err)
	}

	s.FailoverDirection = decoded.FailoverDirection
	s.SourceSiteOperations = decoded.SourceSiteOperations

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ClusterUnplannedFailoverInputProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["providerSpecificDetails"]; ok {
		impl, err := unmarshalClusterUnplannedFailoverProviderSpecificInputImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ProviderSpecificDetails' for 'ClusterUnplannedFailoverInputProperties': %+v", err)
		}
		s.ProviderSpecificDetails = impl
	}
	return nil
}
