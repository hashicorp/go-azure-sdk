package media

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ InputDefinition = FromEachInputFile{}

type FromEachInputFile struct {

	// Fields inherited from InputDefinition
	IncludedTracks *[]TrackDescriptor `json:"includedTracks,omitempty"`
}

var _ json.Marshaler = FromEachInputFile{}

func (s FromEachInputFile) MarshalJSON() ([]byte, error) {
	type wrapper FromEachInputFile
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling FromEachInputFile: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling FromEachInputFile: %+v", err)
	}
	decoded["@odata.type"] = "#Microsoft.Media.FromEachInputFile"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling FromEachInputFile: %+v", err)
	}

	return encoded, nil
}
