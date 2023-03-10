package replicationjobs

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupTaskDetails interface {
}

func unmarshalGroupTaskDetailsImplementation(input []byte) (GroupTaskDetails, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling GroupTaskDetails into map[string]interface: %+v", err)
	}

	value, ok := temp["instanceType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "InlineWorkflowTaskDetails") {
		var out InlineWorkflowTaskDetails
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into InlineWorkflowTaskDetails: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "RecoveryPlanGroupTaskDetails") {
		var out RecoveryPlanGroupTaskDetails
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RecoveryPlanGroupTaskDetails: %+v", err)
		}
		return out, nil
	}

	type RawGroupTaskDetailsImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawGroupTaskDetailsImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
