package subassessments

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdditionalData interface {
}

func unmarshalAdditionalDataImplementation(input []byte) (AdditionalData, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AdditionalData into map[string]interface: %+v", err)
	}

	value, ok := temp["assessedResourceType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "ContainerRegistryVulnerability") {
		var out ContainerRegistryVulnerabilityProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ContainerRegistryVulnerabilityProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "ServerVulnerabilityAssessment") {
		var out ServerVulnerabilityProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ServerVulnerabilityProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "SqlServerVulnerability") {
		var out SqlServerVulnerabilityProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SqlServerVulnerabilityProperties: %+v", err)
		}
		return out, nil
	}

	type RawAdditionalDataImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawAdditionalDataImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
