package settings

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Setting = DataExportSetting{}

type DataExportSetting struct {
	Properties *DataExportSettingProperties `json:"properties,omitempty"`

	// Fields inherited from Setting
	Id   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

var _ json.Marshaler = DataExportSetting{}

func (s DataExportSetting) MarshalJSON() ([]byte, error) {
	type wrapper DataExportSetting
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DataExportSetting: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DataExportSetting: %+v", err)
	}
	decoded["kind"] = "DataExportSetting"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DataExportSetting: %+v", err)
	}

	return encoded, nil
}
