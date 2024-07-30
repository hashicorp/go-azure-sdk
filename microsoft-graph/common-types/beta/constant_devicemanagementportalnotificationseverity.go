package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementPortalNotificationSeverity string

const (
	DeviceManagementPortalNotificationSeverity_Critical      DeviceManagementPortalNotificationSeverity = "critical"
	DeviceManagementPortalNotificationSeverity_Informational DeviceManagementPortalNotificationSeverity = "informational"
	DeviceManagementPortalNotificationSeverity_Unknown       DeviceManagementPortalNotificationSeverity = "unknown"
	DeviceManagementPortalNotificationSeverity_Warning       DeviceManagementPortalNotificationSeverity = "warning"
)

func PossibleValuesForDeviceManagementPortalNotificationSeverity() []string {
	return []string{
		string(DeviceManagementPortalNotificationSeverity_Critical),
		string(DeviceManagementPortalNotificationSeverity_Informational),
		string(DeviceManagementPortalNotificationSeverity_Unknown),
		string(DeviceManagementPortalNotificationSeverity_Warning),
	}
}

func (s *DeviceManagementPortalNotificationSeverity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementPortalNotificationSeverity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementPortalNotificationSeverity(input string) (*DeviceManagementPortalNotificationSeverity, error) {
	vals := map[string]DeviceManagementPortalNotificationSeverity{
		"critical":      DeviceManagementPortalNotificationSeverity_Critical,
		"informational": DeviceManagementPortalNotificationSeverity_Informational,
		"unknown":       DeviceManagementPortalNotificationSeverity_Unknown,
		"warning":       DeviceManagementPortalNotificationSeverity_Warning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementPortalNotificationSeverity(input)
	return &out, nil
}
