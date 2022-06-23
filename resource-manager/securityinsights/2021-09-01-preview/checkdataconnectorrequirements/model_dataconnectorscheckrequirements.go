package checkdataconnectorrequirements

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataConnectorsCheckRequirements interface {
}

func unmarshalDataConnectorsCheckRequirementsImplementation(input []byte) (DataConnectorsCheckRequirements, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DataConnectorsCheckRequirements into map[string]interface: %+v", err)
	}

	value, ok := temp["kind"].(string)
	if !ok {
		return nil, nil
	}

	type RawDataConnectorsCheckRequirementsImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawDataConnectorsCheckRequirementsImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
