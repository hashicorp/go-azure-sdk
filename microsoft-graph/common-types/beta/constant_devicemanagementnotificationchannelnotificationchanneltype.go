package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementNotificationChannelNotificationChannelType string

const (
	DeviceManagementNotificationChannelNotificationChannelType_Email     DeviceManagementNotificationChannelNotificationChannelType = "email"
	DeviceManagementNotificationChannelNotificationChannelType_PhoneCall DeviceManagementNotificationChannelNotificationChannelType = "phoneCall"
	DeviceManagementNotificationChannelNotificationChannelType_Portal    DeviceManagementNotificationChannelNotificationChannelType = "portal"
	DeviceManagementNotificationChannelNotificationChannelType_Sms       DeviceManagementNotificationChannelNotificationChannelType = "sms"
)

func PossibleValuesForDeviceManagementNotificationChannelNotificationChannelType() []string {
	return []string{
		string(DeviceManagementNotificationChannelNotificationChannelType_Email),
		string(DeviceManagementNotificationChannelNotificationChannelType_PhoneCall),
		string(DeviceManagementNotificationChannelNotificationChannelType_Portal),
		string(DeviceManagementNotificationChannelNotificationChannelType_Sms),
	}
}

func (s *DeviceManagementNotificationChannelNotificationChannelType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementNotificationChannelNotificationChannelType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementNotificationChannelNotificationChannelType(input string) (*DeviceManagementNotificationChannelNotificationChannelType, error) {
	vals := map[string]DeviceManagementNotificationChannelNotificationChannelType{
		"email":     DeviceManagementNotificationChannelNotificationChannelType_Email,
		"phonecall": DeviceManagementNotificationChannelNotificationChannelType_PhoneCall,
		"portal":    DeviceManagementNotificationChannelNotificationChannelType_Portal,
		"sms":       DeviceManagementNotificationChannelNotificationChannelType_Sms,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementNotificationChannelNotificationChannelType(input)
	return &out, nil
}
