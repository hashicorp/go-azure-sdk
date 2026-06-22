package replicationevents

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventProviderSpecificDetails interface {
	EventProviderSpecificDetails() BaseEventProviderSpecificDetailsImpl
}

var _ EventProviderSpecificDetails = BaseEventProviderSpecificDetailsImpl{}

type BaseEventProviderSpecificDetailsImpl struct {
	InstanceType string `json:"instanceType"`
}

func (s BaseEventProviderSpecificDetailsImpl) EventProviderSpecificDetails() BaseEventProviderSpecificDetailsImpl {
	return s
}

var _ EventProviderSpecificDetails = RawEventProviderSpecificDetailsImpl{}

// RawEventProviderSpecificDetailsImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawEventProviderSpecificDetailsImpl struct {
	eventProviderSpecificDetails BaseEventProviderSpecificDetailsImpl
	Type                         string
	Values                       map[string]interface{}
}

func (s RawEventProviderSpecificDetailsImpl) EventProviderSpecificDetails() BaseEventProviderSpecificDetailsImpl {
	return s.eventProviderSpecificDetails
}

func (s RawEventProviderSpecificDetailsImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalEventProviderSpecificDetailsImplementation(input []byte) (EventProviderSpecificDetails, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling EventProviderSpecificDetails into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["instanceType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "A2A") {
		var out A2AEventDetails
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into A2AEventDetails: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "HyperVReplica2012") {
		var out HyperVReplica2012EventDetails
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into HyperVReplica2012EventDetails: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "HyperVReplica2012R2") {
		var out HyperVReplica2012R2EventDetails
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into HyperVReplica2012R2EventDetails: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "HyperVReplicaAzure") {
		var out HyperVReplicaAzureEventDetails
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into HyperVReplicaAzureEventDetails: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "HyperVReplicaBaseEventDetails") {
		var out HyperVReplicaBaseEventDetails
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into HyperVReplicaBaseEventDetails: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "InMageAzureV2") {
		var out InMageAzureV2EventDetails
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into InMageAzureV2EventDetails: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "InMageRcm") {
		var out InMageRcmEventDetails
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into InMageRcmEventDetails: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "InMageRcmFailback") {
		var out InMageRcmFailbackEventDetails
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into InMageRcmFailbackEventDetails: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "VMwareCbt") {
		var out VMwareCbtEventDetails
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VMwareCbtEventDetails: %+v", err)
		}
		return out, nil
	}

	var parent BaseEventProviderSpecificDetailsImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseEventProviderSpecificDetailsImpl: %+v", err)
	}

	return RawEventProviderSpecificDetailsImpl{
		eventProviderSpecificDetails: parent,
		Type:                         value,
		Values:                       temp,
	}, nil

}
