package usagedetails

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UsageDetail interface {
}

func unmarshalUsageDetailImplementation(input []byte) (UsageDetail, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling UsageDetail into map[string]interface: %+v", err)
	}

	value, ok := temp["kind"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "legacy") {
		var out LegacyUsageDetail
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LegacyUsageDetail: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "modern") {
		var out ModernUsageDetail
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ModernUsageDetail: %+v", err)
		}
		return out, nil
	}

	type RawUsageDetailImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawUsageDetailImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
