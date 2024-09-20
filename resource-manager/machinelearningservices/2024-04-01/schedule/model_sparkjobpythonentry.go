package schedule

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SparkJobEntry = SparkJobPythonEntry{}

type SparkJobPythonEntry struct {
	File string `json:"file"`

	// Fields inherited from SparkJobEntry

	SparkJobEntryType SparkJobEntryType `json:"sparkJobEntryType"`
}

func (s SparkJobPythonEntry) SparkJobEntry() BaseSparkJobEntryImpl {
	return BaseSparkJobEntryImpl{
		SparkJobEntryType: s.SparkJobEntryType,
	}
}

var _ json.Marshaler = SparkJobPythonEntry{}

func (s SparkJobPythonEntry) MarshalJSON() ([]byte, error) {
	type wrapper SparkJobPythonEntry
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SparkJobPythonEntry: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SparkJobPythonEntry: %+v", err)
	}

	decoded["sparkJobEntryType"] = "SparkJobPythonEntry"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SparkJobPythonEntry: %+v", err)
	}

	return encoded, nil
}
