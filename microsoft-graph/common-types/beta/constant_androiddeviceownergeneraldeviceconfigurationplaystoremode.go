package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerGeneralDeviceConfigurationPlayStoreMode string

const (
	AndroidDeviceOwnerGeneralDeviceConfigurationPlayStoreMode_AllowList     AndroidDeviceOwnerGeneralDeviceConfigurationPlayStoreMode = "allowList"
	AndroidDeviceOwnerGeneralDeviceConfigurationPlayStoreMode_BlockList     AndroidDeviceOwnerGeneralDeviceConfigurationPlayStoreMode = "blockList"
	AndroidDeviceOwnerGeneralDeviceConfigurationPlayStoreMode_NotConfigured AndroidDeviceOwnerGeneralDeviceConfigurationPlayStoreMode = "notConfigured"
)

func PossibleValuesForAndroidDeviceOwnerGeneralDeviceConfigurationPlayStoreMode() []string {
	return []string{
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPlayStoreMode_AllowList),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPlayStoreMode_BlockList),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPlayStoreMode_NotConfigured),
	}
}

func (s *AndroidDeviceOwnerGeneralDeviceConfigurationPlayStoreMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerGeneralDeviceConfigurationPlayStoreMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerGeneralDeviceConfigurationPlayStoreMode(input string) (*AndroidDeviceOwnerGeneralDeviceConfigurationPlayStoreMode, error) {
	vals := map[string]AndroidDeviceOwnerGeneralDeviceConfigurationPlayStoreMode{
		"allowlist":     AndroidDeviceOwnerGeneralDeviceConfigurationPlayStoreMode_AllowList,
		"blocklist":     AndroidDeviceOwnerGeneralDeviceConfigurationPlayStoreMode_BlockList,
		"notconfigured": AndroidDeviceOwnerGeneralDeviceConfigurationPlayStoreMode_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerGeneralDeviceConfigurationPlayStoreMode(input)
	return &out, nil
}
