package alertsmanagements

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AlertsMetaDataProperties = MonitorServiceList{}

type MonitorServiceList struct {
	Data []MonitorServiceDetails `json:"data"`

	// Fields inherited from AlertsMetaDataProperties

	MetadataIdentifier MetadataIdentifier `json:"metadataIdentifier"`
}

func (s MonitorServiceList) AlertsMetaDataProperties() BaseAlertsMetaDataPropertiesImpl {
	return BaseAlertsMetaDataPropertiesImpl{
		MetadataIdentifier: s.MetadataIdentifier,
	}
}

var _ json.Marshaler = MonitorServiceList{}

func (s MonitorServiceList) MarshalJSON() ([]byte, error) {
	type wrapper MonitorServiceList
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MonitorServiceList: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MonitorServiceList: %+v", err)
	}

	decoded["metadataIdentifier"] = "MonitorServiceList"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MonitorServiceList: %+v", err)
	}

	return encoded, nil
}
