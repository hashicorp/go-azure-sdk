package iotcentrals

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ JobScheduleEnd = DateJobScheduleEnd{}

type DateJobScheduleEnd struct {
	Date string `json:"date"`

	// Fields inherited from JobScheduleEnd

	Type string `json:"type"`
}

func (s DateJobScheduleEnd) JobScheduleEnd() BaseJobScheduleEndImpl {
	return BaseJobScheduleEndImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = DateJobScheduleEnd{}

func (s DateJobScheduleEnd) MarshalJSON() ([]byte, error) {
	type wrapper DateJobScheduleEnd
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DateJobScheduleEnd: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DateJobScheduleEnd: %+v", err)
	}

	decoded["type"] = "date"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DateJobScheduleEnd: %+v", err)
	}

	return encoded, nil
}
