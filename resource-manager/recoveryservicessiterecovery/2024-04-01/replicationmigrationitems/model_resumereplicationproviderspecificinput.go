package replicationmigrationitems

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResumeReplicationProviderSpecificInput interface {
	ResumeReplicationProviderSpecificInput() BaseResumeReplicationProviderSpecificInputImpl
}

var _ ResumeReplicationProviderSpecificInput = BaseResumeReplicationProviderSpecificInputImpl{}

type BaseResumeReplicationProviderSpecificInputImpl struct {
	InstanceType string `json:"instanceType"`
}

func (s BaseResumeReplicationProviderSpecificInputImpl) ResumeReplicationProviderSpecificInput() BaseResumeReplicationProviderSpecificInputImpl {
	return s
}

var _ ResumeReplicationProviderSpecificInput = RawResumeReplicationProviderSpecificInputImpl{}

// RawResumeReplicationProviderSpecificInputImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawResumeReplicationProviderSpecificInputImpl struct {
	resumeReplicationProviderSpecificInput BaseResumeReplicationProviderSpecificInputImpl
	Type                                   string
	Values                                 map[string]interface{}
}

func (s RawResumeReplicationProviderSpecificInputImpl) ResumeReplicationProviderSpecificInput() BaseResumeReplicationProviderSpecificInputImpl {
	return s.resumeReplicationProviderSpecificInput
}

func (s RawResumeReplicationProviderSpecificInputImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalResumeReplicationProviderSpecificInputImplementation(input []byte) (ResumeReplicationProviderSpecificInput, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ResumeReplicationProviderSpecificInput into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["instanceType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "VMwareCbt") {
		var out VMwareCbtResumeReplicationInput
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VMwareCbtResumeReplicationInput: %+v", err)
		}
		return out, nil
	}

	var parent BaseResumeReplicationProviderSpecificInputImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseResumeReplicationProviderSpecificInputImpl: %+v", err)
	}

	return RawResumeReplicationProviderSpecificInputImpl{
		resumeReplicationProviderSpecificInput: parent,
		Type:                                   value,
		Values:                                 temp,
	}, nil

}
