package schedule

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MonitoringAlertNotificationSettingsBase = AzMonMonitoringAlertNotificationSettings{}

type AzMonMonitoringAlertNotificationSettings struct {

	// Fields inherited from MonitoringAlertNotificationSettingsBase
}

var _ json.Marshaler = AzMonMonitoringAlertNotificationSettings{}

func (s AzMonMonitoringAlertNotificationSettings) MarshalJSON() ([]byte, error) {
	type wrapper AzMonMonitoringAlertNotificationSettings
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzMonMonitoringAlertNotificationSettings: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzMonMonitoringAlertNotificationSettings: %+v", err)
	}
	decoded["alertNotificationType"] = "AzureMonitor"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzMonMonitoringAlertNotificationSettings: %+v", err)
	}

	return encoded, nil
}
