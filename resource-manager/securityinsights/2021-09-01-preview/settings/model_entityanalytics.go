package settings

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Settings = EntityAnalytics{}

type EntityAnalytics struct {
	Properties *EntityAnalyticsProperties `json:"properties,omitempty"`

	// Fields inherited from Settings
	Etag       *string                `json:"etag,omitempty"`
	Id         *string                `json:"id,omitempty"`
	Name       *string                `json:"name,omitempty"`
	SystemData *systemdata.SystemData `json:"systemData,omitempty"`
	Type       *string                `json:"type,omitempty"`
}

var _ json.Marshaler = EntityAnalytics{}

func (s EntityAnalytics) MarshalJSON() ([]byte, error) {
	type wrapper EntityAnalytics
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EntityAnalytics: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EntityAnalytics: %+v", err)
	}
	decoded["kind"] = "EntityAnalytics"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EntityAnalytics: %+v", err)
	}

	return encoded, nil
}
