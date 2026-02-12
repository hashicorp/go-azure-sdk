package iotcentrals

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ JobScheduleEnd = OccurrencesJobScheduleEnd{}

type OccurrencesJobScheduleEnd struct {
	Occurrences int64 `json:"occurrences"`

	// Fields inherited from JobScheduleEnd

	Type string `json:"type"`
}

func (s OccurrencesJobScheduleEnd) JobScheduleEnd() BaseJobScheduleEndImpl {
	return BaseJobScheduleEndImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = OccurrencesJobScheduleEnd{}

func (s OccurrencesJobScheduleEnd) MarshalJSON() ([]byte, error) {
	type wrapper OccurrencesJobScheduleEnd
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OccurrencesJobScheduleEnd: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OccurrencesJobScheduleEnd: %+v", err)
	}

	decoded["type"] = "occurrences"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OccurrencesJobScheduleEnd: %+v", err)
	}

	return encoded, nil
}
