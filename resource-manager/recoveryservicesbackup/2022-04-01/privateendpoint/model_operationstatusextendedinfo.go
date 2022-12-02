package privateendpoint

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OperationStatusExtendedInfo interface {
}

func unmarshalOperationStatusExtendedInfoImplementation(input []byte) (OperationStatusExtendedInfo, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling OperationStatusExtendedInfo into map[string]interface: %+v", err)
	}

	value, ok := temp["objectType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "OperationStatusJobExtendedInfo") {
		var out OperationStatusJobExtendedInfo
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OperationStatusJobExtendedInfo: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "OperationStatusJobsExtendedInfo") {
		var out OperationStatusJobsExtendedInfo
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OperationStatusJobsExtendedInfo: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "OperationStatusProvisionILRExtendedInfo") {
		var out OperationStatusProvisionILRExtendedInfo
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OperationStatusProvisionILRExtendedInfo: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "OperationStatusValidateOperationExtendedInfo") {
		var out OperationStatusValidateOperationExtendedInfo
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OperationStatusValidateOperationExtendedInfo: %+v", err)
		}
		return out, nil
	}

	type RawOperationStatusExtendedInfoImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawOperationStatusExtendedInfoImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
