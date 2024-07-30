package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequireUnlock string

const (
	AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequireUnlock_Daily             AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequireUnlock = "daily"
	AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequireUnlock_DeviceDefault     AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequireUnlock = "deviceDefault"
	AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequireUnlock_UnkownFutureValue AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequireUnlock = "unkownFutureValue"
)

func PossibleValuesForAndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequireUnlock() []string {
	return []string{
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequireUnlock_Daily),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequireUnlock_DeviceDefault),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequireUnlock_UnkownFutureValue),
	}
}

func (s *AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequireUnlock) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequireUnlock(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequireUnlock(input string) (*AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequireUnlock, error) {
	vals := map[string]AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequireUnlock{
		"daily":             AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequireUnlock_Daily,
		"devicedefault":     AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequireUnlock_DeviceDefault,
		"unkownfuturevalue": AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequireUnlock_UnkownFutureValue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequireUnlock(input)
	return &out, nil
}
