package entitytypes

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UrlEntity{}

type UrlEntity struct {
	Properties *UrlEntityProperties `json:"properties,omitempty"`

	// Fields inherited from Entity
	Id         *string     `json:"id,omitempty"`
	Name       *string     `json:"name,omitempty"`
	SystemData *SystemData `json:"systemData,omitempty"`
	Type       *string     `json:"type,omitempty"`
}

var _ json.Marshaler = UrlEntity{}

func (s UrlEntity) MarshalJSON() ([]byte, error) {
	type wrapper UrlEntity
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UrlEntity: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UrlEntity: %+v", err)
	}
	decoded["kind"] = "Url"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UrlEntity: %+v", err)
	}

	return encoded, nil
}
