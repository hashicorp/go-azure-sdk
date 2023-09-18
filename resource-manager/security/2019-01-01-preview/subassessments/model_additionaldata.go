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

// RawAdditionalDataImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAdditionalDataImpl struct {
	Type   string
	Values map[string]interface{}
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

	out := RawAdditionalDataImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
