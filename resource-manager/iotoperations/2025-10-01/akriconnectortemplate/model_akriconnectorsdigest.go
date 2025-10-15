package akriconnectortemplate

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AkriConnectorsTagDigestSettings = AkriConnectorsDigest{}

type AkriConnectorsDigest struct {
	Digest string `json:"digest"`

	// Fields inherited from AkriConnectorsTagDigestSettings

	TagDigestType AkriConnectorsTagDigestType `json:"tagDigestType"`
}

func (s AkriConnectorsDigest) AkriConnectorsTagDigestSettings() BaseAkriConnectorsTagDigestSettingsImpl {
	return BaseAkriConnectorsTagDigestSettingsImpl{
		TagDigestType: s.TagDigestType,
	}
}

var _ json.Marshaler = AkriConnectorsDigest{}

func (s AkriConnectorsDigest) MarshalJSON() ([]byte, error) {
	type wrapper AkriConnectorsDigest
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AkriConnectorsDigest: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AkriConnectorsDigest: %+v", err)
	}

	decoded["tagDigestType"] = "Digest"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AkriConnectorsDigest: %+v", err)
	}

	return encoded, nil
}
