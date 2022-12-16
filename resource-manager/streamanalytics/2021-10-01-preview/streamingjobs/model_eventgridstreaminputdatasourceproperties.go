package streamingjobs

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventGridStreamInputDataSourceProperties struct {
	EventTypes      *[]string                 `json:"eventTypes,omitempty"`
	Schema          *EventGridEventSchemaType `json:"schema,omitempty"`
	StorageAccounts *[]StorageAccount         `json:"storageAccounts,omitempty"`
	Subscriber      StreamInputDataSource     `json:"subscriber"`
}

var _ json.Unmarshaler = &EventGridStreamInputDataSourceProperties{}

func (s *EventGridStreamInputDataSourceProperties) UnmarshalJSON(bytes []byte) error {
	type alias EventGridStreamInputDataSourceProperties
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into EventGridStreamInputDataSourceProperties: %+v", err)
	}

	s.EventTypes = decoded.EventTypes
	s.Schema = decoded.Schema
	s.StorageAccounts = decoded.StorageAccounts

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling EventGridStreamInputDataSourceProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["subscriber"]; ok {
		impl, err := unmarshalStreamInputDataSourceImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Subscriber' for 'EventGridStreamInputDataSourceProperties': %+v", err)
		}
		s.Subscriber = impl
	}
	return nil
}
