package iotcentrals

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Destination interface {
	Destination() BaseDestinationImpl
}

var _ Destination = BaseDestinationImpl{}

type BaseDestinationImpl struct {
	DisplayName    string             `json:"displayName"`
	Errors         *[]DataExportError `json:"errors,omitempty"`
	Id             *string            `json:"id,omitempty"`
	LastExportTime *string            `json:"lastExportTime,omitempty"`
	Status         *string            `json:"status,omitempty"`
	Type           string             `json:"type"`
}

func (s BaseDestinationImpl) Destination() BaseDestinationImpl {
	return s
}

var _ Destination = RawDestinationImpl{}

// RawDestinationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDestinationImpl struct {
	destination BaseDestinationImpl
	Type        string
	Values      map[string]interface{}
}

func (s RawDestinationImpl) Destination() BaseDestinationImpl {
	return s.destination
}

func UnmarshalDestinationImplementation(input []byte) (Destination, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling Destination into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "blobstorage@v1") {
		var out BlobStorageV1Destination
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BlobStorageV1Destination: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "dataexplorer@v1") {
		var out DataExplorerV1Destination
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DataExplorerV1Destination: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "eventhubs@v1") {
		var out EventHubsV1Destination
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EventHubsV1Destination: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "ExportDestination") {
		var out ExportDestination
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExportDestination: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "servicebusqueue@v1") {
		var out ServiceBusQueueV1Destination
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ServiceBusQueueV1Destination: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "servicebustopic@v1") {
		var out ServiceBusTopicV1Destination
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ServiceBusTopicV1Destination: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "webhook@v1") {
		var out WebhookV1Destination
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WebhookV1Destination: %+v", err)
		}
		return out, nil
	}

	var parent BaseDestinationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDestinationImpl: %+v", err)
	}

	return RawDestinationImpl{
		destination: parent,
		Type:        value,
		Values:      temp,
	}, nil

}
