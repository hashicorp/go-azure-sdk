package threatintelligence

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ThreatIntelligenceInformation interface {
}

func unmarshalThreatIntelligenceInformationImplementation(input []byte) (ThreatIntelligenceInformation, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ThreatIntelligenceInformation into map[string]interface: %+v", err)
	}

	value, ok := temp["kind"].(string)
	if !ok {
		return nil, nil
	}

	type RawThreatIntelligenceInformationImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawThreatIntelligenceInformationImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
