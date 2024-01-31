package connecttosourcesqlservertasks

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectToSourceSqlServerTaskOutput interface {
}

// RawConnectToSourceSqlServerTaskOutputImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawConnectToSourceSqlServerTaskOutputImpl struct {
	Type   string
	Values map[string]interface{}
}

func unmarshalConnectToSourceSqlServerTaskOutputImplementation(input []byte) (ConnectToSourceSqlServerTaskOutput, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ConnectToSourceSqlServerTaskOutput into map[string]interface: %+v", err)
	}

	value, ok := temp["resultType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "AgentJobLevelOutput") {
		var out ConnectToSourceSqlServerTaskOutputAgentJobLevel
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ConnectToSourceSqlServerTaskOutputAgentJobLevel: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "DatabaseLevelOutput") {
		var out ConnectToSourceSqlServerTaskOutputDatabaseLevel
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ConnectToSourceSqlServerTaskOutputDatabaseLevel: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "LoginLevelOutput") {
		var out ConnectToSourceSqlServerTaskOutputLoginLevel
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ConnectToSourceSqlServerTaskOutputLoginLevel: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "TaskLevelOutput") {
		var out ConnectToSourceSqlServerTaskOutputTaskLevel
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ConnectToSourceSqlServerTaskOutputTaskLevel: %+v", err)
		}
		return out, nil
	}

	out := RawConnectToSourceSqlServerTaskOutputImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
