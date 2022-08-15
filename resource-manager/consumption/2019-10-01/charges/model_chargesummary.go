package charges

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChargeSummary interface {
}

func unmarshalChargeSummaryImplementation(input []byte) (ChargeSummary, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ChargeSummary into map[string]interface: %+v", err)
	}

	value, ok := temp["kind"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "legacy") {
		var out LegacyChargeSummary
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LegacyChargeSummary: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "modern") {
		var out ModernChargeSummary
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ModernChargeSummary: %+v", err)
		}
		return out, nil
	}

	type RawChargeSummaryImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawChargeSummaryImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
