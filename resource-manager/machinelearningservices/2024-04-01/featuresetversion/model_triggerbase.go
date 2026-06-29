package featuresetversion

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TriggerBase interface {
	TriggerBase() BaseTriggerBaseImpl
}

var _ TriggerBase = BaseTriggerBaseImpl{}

type BaseTriggerBaseImpl struct {
	EndTime     *string     `json:"endTime,omitempty"`
	StartTime   *string     `json:"startTime,omitempty"`
	TimeZone    *string     `json:"timeZone,omitempty"`
	TriggerType TriggerType `json:"triggerType"`
}

func (s BaseTriggerBaseImpl) TriggerBase() BaseTriggerBaseImpl {
	return s
}

var _ TriggerBase = RawTriggerBaseImpl{}

// RawTriggerBaseImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawTriggerBaseImpl struct {
	triggerBase BaseTriggerBaseImpl
	Type        string
	Values      map[string]interface{}
}

func (s RawTriggerBaseImpl) TriggerBase() BaseTriggerBaseImpl {
	return s.triggerBase
}

func (s RawTriggerBaseImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalTriggerBaseImplementation(input []byte) (TriggerBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling TriggerBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["triggerType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "Cron") {
		var out CronTrigger
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CronTrigger: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Recurrence") {
		var out RecurrenceTrigger
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RecurrenceTrigger: %+v", err)
		}
		return out, nil
	}

	var parent BaseTriggerBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseTriggerBaseImpl: %+v", err)
	}

	return RawTriggerBaseImpl{
		triggerBase: parent,
		Type:        value,
		Values:      temp,
	}, nil

}
