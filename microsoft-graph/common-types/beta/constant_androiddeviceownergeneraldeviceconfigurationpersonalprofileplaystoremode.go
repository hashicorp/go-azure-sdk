package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerGeneralDeviceConfigurationPersonalProfilePlayStoreMode string

const (
	AndroidDeviceOwnerGeneralDeviceConfigurationPersonalProfilePlayStoreMode_AllowedApps   AndroidDeviceOwnerGeneralDeviceConfigurationPersonalProfilePlayStoreMode = "allowedApps"
	AndroidDeviceOwnerGeneralDeviceConfigurationPersonalProfilePlayStoreMode_BlockedApps   AndroidDeviceOwnerGeneralDeviceConfigurationPersonalProfilePlayStoreMode = "blockedApps"
	AndroidDeviceOwnerGeneralDeviceConfigurationPersonalProfilePlayStoreMode_NotConfigured AndroidDeviceOwnerGeneralDeviceConfigurationPersonalProfilePlayStoreMode = "notConfigured"
)

func PossibleValuesForAndroidDeviceOwnerGeneralDeviceConfigurationPersonalProfilePlayStoreMode() []string {
	return []string{
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPersonalProfilePlayStoreMode_AllowedApps),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPersonalProfilePlayStoreMode_BlockedApps),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPersonalProfilePlayStoreMode_NotConfigured),
	}
}

func (s *AndroidDeviceOwnerGeneralDeviceConfigurationPersonalProfilePlayStoreMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerGeneralDeviceConfigurationPersonalProfilePlayStoreMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerGeneralDeviceConfigurationPersonalProfilePlayStoreMode(input string) (*AndroidDeviceOwnerGeneralDeviceConfigurationPersonalProfilePlayStoreMode, error) {
	vals := map[string]AndroidDeviceOwnerGeneralDeviceConfigurationPersonalProfilePlayStoreMode{
		"allowedapps":   AndroidDeviceOwnerGeneralDeviceConfigurationPersonalProfilePlayStoreMode_AllowedApps,
		"blockedapps":   AndroidDeviceOwnerGeneralDeviceConfigurationPersonalProfilePlayStoreMode_BlockedApps,
		"notconfigured": AndroidDeviceOwnerGeneralDeviceConfigurationPersonalProfilePlayStoreMode_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerGeneralDeviceConfigurationPersonalProfilePlayStoreMode(input)
	return &out, nil
}
