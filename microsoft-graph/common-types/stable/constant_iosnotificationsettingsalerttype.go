package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosNotificationSettingsAlertType string

const (
	IosNotificationSettingsAlertType_Banner        IosNotificationSettingsAlertType = "banner"
	IosNotificationSettingsAlertType_DeviceDefault IosNotificationSettingsAlertType = "deviceDefault"
	IosNotificationSettingsAlertType_Modal         IosNotificationSettingsAlertType = "modal"
	IosNotificationSettingsAlertType_None          IosNotificationSettingsAlertType = "none"
)

func PossibleValuesForIosNotificationSettingsAlertType() []string {
	return []string{
		string(IosNotificationSettingsAlertType_Banner),
		string(IosNotificationSettingsAlertType_DeviceDefault),
		string(IosNotificationSettingsAlertType_Modal),
		string(IosNotificationSettingsAlertType_None),
	}
}

func (s *IosNotificationSettingsAlertType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosNotificationSettingsAlertType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosNotificationSettingsAlertType(input string) (*IosNotificationSettingsAlertType, error) {
	vals := map[string]IosNotificationSettingsAlertType{
		"banner":        IosNotificationSettingsAlertType_Banner,
		"devicedefault": IosNotificationSettingsAlertType_DeviceDefault,
		"modal":         IosNotificationSettingsAlertType_Modal,
		"none":          IosNotificationSettingsAlertType_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosNotificationSettingsAlertType(input)
	return &out, nil
}
