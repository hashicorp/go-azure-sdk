package extendedueinformation

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ExtendedUeInfoProperties = UeInfo4G{}

type UeInfo4G struct {
	Info UeInfo4GProperties `json:"info"`

	// Fields inherited from ExtendedUeInfoProperties
	LastReadAt *string `json:"lastReadAt,omitempty"`
}

var _ json.Marshaler = UeInfo4G{}

func (s UeInfo4G) MarshalJSON() ([]byte, error) {
	type wrapper UeInfo4G
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UeInfo4G: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UeInfo4G: %+v", err)
	}
	decoded["ratType"] = "4G"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UeInfo4G: %+v", err)
	}

	return encoded, nil
}
