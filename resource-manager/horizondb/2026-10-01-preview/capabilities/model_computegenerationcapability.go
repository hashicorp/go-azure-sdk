package capabilities

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComputeGenerationCapability struct {
	ComputeGeneration           int64                    `json:"computeGeneration"`
	RecommendedComputeModelType PoolComputeModelType     `json:"recommendedComputeModelType"`
	SupportedComputeModels      []ComputeModelCapability `json:"supportedComputeModels"`
}

var _ json.Unmarshaler = &ComputeGenerationCapability{}

func (s *ComputeGenerationCapability) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ComputeGeneration           int64                `json:"computeGeneration"`
		RecommendedComputeModelType PoolComputeModelType `json:"recommendedComputeModelType"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ComputeGeneration = decoded.ComputeGeneration
	s.RecommendedComputeModelType = decoded.RecommendedComputeModelType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ComputeGenerationCapability into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["supportedComputeModels"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling SupportedComputeModels into list []json.RawMessage: %+v", err)
		}

		output := make([]ComputeModelCapability, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalComputeModelCapabilityImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'SupportedComputeModels' for 'ComputeGenerationCapability': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.SupportedComputeModels = output
	}

	return nil
}
