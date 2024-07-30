package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequireUnlock string

const (
	AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequireUnlock_Daily             AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequireUnlock = "daily"
	AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequireUnlock_DeviceDefault     AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequireUnlock = "deviceDefault"
	AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequireUnlock_UnkownFutureValue AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequireUnlock = "unkownFutureValue"
)

func PossibleValuesForAndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequireUnlock() []string {
	return []string{
		string(AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequireUnlock_Daily),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequireUnlock_DeviceDefault),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequireUnlock_UnkownFutureValue),
	}
}

func (s *AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequireUnlock) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequireUnlock(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequireUnlock(input string) (*AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequireUnlock, error) {
	vals := map[string]AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequireUnlock{
		"daily":             AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequireUnlock_Daily,
		"devicedefault":     AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequireUnlock_DeviceDefault,
		"unkownfuturevalue": AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequireUnlock_UnkownFutureValue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequireUnlock(input)
	return &out, nil
}
