package schedule

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MonitoringAlertNotificationSettingsBase = EmailMonitoringAlertNotificationSettings{}

type EmailMonitoringAlertNotificationSettings struct {
	EmailNotificationSetting *NotificationSetting `json:"emailNotificationSetting,omitempty"`

	// Fields inherited from MonitoringAlertNotificationSettingsBase
}

var _ json.Marshaler = EmailMonitoringAlertNotificationSettings{}

func (s EmailMonitoringAlertNotificationSettings) MarshalJSON() ([]byte, error) {
	type wrapper EmailMonitoringAlertNotificationSettings
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EmailMonitoringAlertNotificationSettings: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EmailMonitoringAlertNotificationSettings: %+v", err)
	}
	decoded["alertNotificationType"] = "Email"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EmailMonitoringAlertNotificationSettings: %+v", err)
	}

	return encoded, nil
}
