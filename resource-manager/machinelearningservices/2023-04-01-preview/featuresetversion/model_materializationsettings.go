package featuresetversion

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MaterializationSettings struct {
	Notification       *NotificationSetting            `json:"notification,omitempty"`
	Resource           *MaterializationComputeResource `json:"resource,omitempty"`
	Schedule           TriggerBase                     `json:"schedule"`
	SparkConfiguration *map[string]string              `json:"sparkConfiguration,omitempty"`
	StoreType          *MaterializationStoreType       `json:"storeType,omitempty"`
}

var _ json.Unmarshaler = &MaterializationSettings{}

func (s *MaterializationSettings) UnmarshalJSON(bytes []byte) error {
	type alias MaterializationSettings
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into MaterializationSettings: %+v", err)
	}

	s.Notification = decoded.Notification
	s.Resource = decoded.Resource
	s.SparkConfiguration = decoded.SparkConfiguration
	s.StoreType = decoded.StoreType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling MaterializationSettings into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["schedule"]; ok {
		impl, err := unmarshalTriggerBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Schedule' for 'MaterializationSettings': %+v", err)
		}
		s.Schedule = impl
	}
	return nil
}
