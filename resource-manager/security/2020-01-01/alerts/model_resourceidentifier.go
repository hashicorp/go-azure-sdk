package alerts

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceIdentifier interface {
}

func unmarshalResourceIdentifierImplementation(input []byte) (ResourceIdentifier, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ResourceIdentifier into map[string]interface: %+v", err)
	}

	value, ok := temp["type"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "AzureResource") {
		var out AzureResourceIdentifier
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureResourceIdentifier: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "LogAnalytics") {
		var out LogAnalyticsIdentifier
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LogAnalyticsIdentifier: %+v", err)
		}
		return out, nil
	}

	type RawResourceIdentifierImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawResourceIdentifierImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
