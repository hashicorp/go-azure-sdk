package replicationjobs

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ JobDetails = AsrJobDetails{}

type AsrJobDetails struct {

	// Fields inherited from JobDetails
	AffectedObjectDetails *map[string]string `json:"affectedObjectDetails,omitempty"`
}

var _ json.Marshaler = AsrJobDetails{}

func (s AsrJobDetails) MarshalJSON() ([]byte, error) {
	type wrapper AsrJobDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AsrJobDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AsrJobDetails: %+v", err)
	}
	decoded["instanceType"] = "AsrJobDetails"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AsrJobDetails: %+v", err)
	}

	return encoded, nil
}
