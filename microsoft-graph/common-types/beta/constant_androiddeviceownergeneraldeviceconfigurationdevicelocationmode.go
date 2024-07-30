package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerGeneralDeviceConfigurationDeviceLocationMode string

const (
	AndroidDeviceOwnerGeneralDeviceConfigurationDeviceLocationMode_Disabled      AndroidDeviceOwnerGeneralDeviceConfigurationDeviceLocationMode = "disabled"
	AndroidDeviceOwnerGeneralDeviceConfigurationDeviceLocationMode_NotConfigured AndroidDeviceOwnerGeneralDeviceConfigurationDeviceLocationMode = "notConfigured"
)

func PossibleValuesForAndroidDeviceOwnerGeneralDeviceConfigurationDeviceLocationMode() []string {
	return []string{
		string(AndroidDeviceOwnerGeneralDeviceConfigurationDeviceLocationMode_Disabled),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationDeviceLocationMode_NotConfigured),
	}
}

func (s *AndroidDeviceOwnerGeneralDeviceConfigurationDeviceLocationMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerGeneralDeviceConfigurationDeviceLocationMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerGeneralDeviceConfigurationDeviceLocationMode(input string) (*AndroidDeviceOwnerGeneralDeviceConfigurationDeviceLocationMode, error) {
	vals := map[string]AndroidDeviceOwnerGeneralDeviceConfigurationDeviceLocationMode{
		"disabled":      AndroidDeviceOwnerGeneralDeviceConfigurationDeviceLocationMode_Disabled,
		"notconfigured": AndroidDeviceOwnerGeneralDeviceConfigurationDeviceLocationMode_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerGeneralDeviceConfigurationDeviceLocationMode(input)
	return &out, nil
}
