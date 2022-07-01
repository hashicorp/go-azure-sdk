package media

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ InputDefinition = InputFile{}

type InputFile struct {
	Filename *string `json:"filename,omitempty"`

	// Fields inherited from InputDefinition
	IncludedTracks *[]TrackDescriptor `json:"includedTracks,omitempty"`
}

var _ json.Marshaler = InputFile{}

func (s InputFile) MarshalJSON() ([]byte, error) {
	type wrapper InputFile
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling InputFile: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling InputFile: %+v", err)
	}
	decoded["@odata.type"] = "#Microsoft.Media.InputFile"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling InputFile: %+v", err)
	}

	return encoded, nil
}
