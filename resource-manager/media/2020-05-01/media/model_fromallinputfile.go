package media

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ InputDefinition = FromAllInputFile{}

type FromAllInputFile struct {

	// Fields inherited from InputDefinition
	IncludedTracks *[]TrackDescriptor `json:"includedTracks,omitempty"`
}

var _ json.Marshaler = FromAllInputFile{}

func (s FromAllInputFile) MarshalJSON() ([]byte, error) {
	type wrapper FromAllInputFile
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling FromAllInputFile: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling FromAllInputFile: %+v", err)
	}
	decoded["@odata.type"] = "#Microsoft.Media.FromAllInputFile"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling FromAllInputFile: %+v", err)
	}

	return encoded, nil
}
