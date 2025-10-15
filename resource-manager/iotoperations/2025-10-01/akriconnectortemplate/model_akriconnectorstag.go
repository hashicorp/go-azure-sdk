package akriconnectortemplate

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AkriConnectorsTagDigestSettings = AkriConnectorsTag{}

type AkriConnectorsTag struct {
	Tag string `json:"tag"`

	// Fields inherited from AkriConnectorsTagDigestSettings

	TagDigestType AkriConnectorsTagDigestType `json:"tagDigestType"`
}

func (s AkriConnectorsTag) AkriConnectorsTagDigestSettings() BaseAkriConnectorsTagDigestSettingsImpl {
	return BaseAkriConnectorsTagDigestSettingsImpl{
		TagDigestType: s.TagDigestType,
	}
}

var _ json.Marshaler = AkriConnectorsTag{}

func (s AkriConnectorsTag) MarshalJSON() ([]byte, error) {
	type wrapper AkriConnectorsTag
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AkriConnectorsTag: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AkriConnectorsTag: %+v", err)
	}

	decoded["tagDigestType"] = "Tag"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AkriConnectorsTag: %+v", err)
	}

	return encoded, nil
}
