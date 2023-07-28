package job

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SparkJobEntry = SparkJobScalaEntry{}

type SparkJobScalaEntry struct {
	ClassName string `json:"className"`

	// Fields inherited from SparkJobEntry
}

var _ json.Marshaler = SparkJobScalaEntry{}

func (s SparkJobScalaEntry) MarshalJSON() ([]byte, error) {
	type wrapper SparkJobScalaEntry
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SparkJobScalaEntry: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SparkJobScalaEntry: %+v", err)
	}
	decoded["sparkJobEntryType"] = "SparkJobScalaEntry"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SparkJobScalaEntry: %+v", err)
	}

	return encoded, nil
}
