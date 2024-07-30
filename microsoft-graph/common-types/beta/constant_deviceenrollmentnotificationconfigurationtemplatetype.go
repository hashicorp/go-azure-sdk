package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceEnrollmentNotificationConfigurationTemplateType string

const (
	DeviceEnrollmentNotificationConfigurationTemplateType_Email DeviceEnrollmentNotificationConfigurationTemplateType = "email"
	DeviceEnrollmentNotificationConfigurationTemplateType_Push  DeviceEnrollmentNotificationConfigurationTemplateType = "push"
)

func PossibleValuesForDeviceEnrollmentNotificationConfigurationTemplateType() []string {
	return []string{
		string(DeviceEnrollmentNotificationConfigurationTemplateType_Email),
		string(DeviceEnrollmentNotificationConfigurationTemplateType_Push),
	}
}

func (s *DeviceEnrollmentNotificationConfigurationTemplateType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceEnrollmentNotificationConfigurationTemplateType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceEnrollmentNotificationConfigurationTemplateType(input string) (*DeviceEnrollmentNotificationConfigurationTemplateType, error) {
	vals := map[string]DeviceEnrollmentNotificationConfigurationTemplateType{
		"email": DeviceEnrollmentNotificationConfigurationTemplateType_Email,
		"push":  DeviceEnrollmentNotificationConfigurationTemplateType_Push,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceEnrollmentNotificationConfigurationTemplateType(input)
	return &out, nil
}
