package iotcentrals

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobData interface {
	JobData() BaseJobDataImpl
}

var _ JobData = BaseJobDataImpl{}

type BaseJobDataImpl struct {
	Type string `json:"type"`
}

func (s BaseJobDataImpl) JobData() BaseJobDataImpl {
	return s
}

var _ JobData = RawJobDataImpl{}

// RawJobDataImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawJobDataImpl struct {
	jobData BaseJobDataImpl
	Type    string
	Values  map[string]interface{}
}

func (s RawJobDataImpl) JobData() BaseJobDataImpl {
	return s.jobData
}

func UnmarshalJobDataImplementation(input []byte) (JobData, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling JobData into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "cloudProperty") {
		var out CloudPropertyJobData
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPropertyJobData: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "command") {
		var out CommandJobData
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CommandJobData: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "deviceManifestMigration") {
		var out DeviceManifestMigrationJobData
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManifestMigrationJobData: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "deviceTemplateMigration") {
		var out DeviceTemplateMigrationJobData
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceTemplateMigrationJobData: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "property") {
		var out PropertyJobData
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PropertyJobData: %+v", err)
		}
		return out, nil
	}

	var parent BaseJobDataImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseJobDataImpl: %+v", err)
	}

	return RawJobDataImpl{
		jobData: parent,
		Type:    value,
		Values:  temp,
	}, nil

}
