package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationCategoryPlatforms string

const (
	DeviceManagementConfigurationCategoryPlatforms_Android    DeviceManagementConfigurationCategoryPlatforms = "android"
	DeviceManagementConfigurationCategoryPlatforms_IOS        DeviceManagementConfigurationCategoryPlatforms = "iOS"
	DeviceManagementConfigurationCategoryPlatforms_Linux      DeviceManagementConfigurationCategoryPlatforms = "linux"
	DeviceManagementConfigurationCategoryPlatforms_MacOS      DeviceManagementConfigurationCategoryPlatforms = "macOS"
	DeviceManagementConfigurationCategoryPlatforms_None       DeviceManagementConfigurationCategoryPlatforms = "none"
	DeviceManagementConfigurationCategoryPlatforms_Windows10  DeviceManagementConfigurationCategoryPlatforms = "windows10"
	DeviceManagementConfigurationCategoryPlatforms_Windows10X DeviceManagementConfigurationCategoryPlatforms = "windows10X"
)

func PossibleValuesForDeviceManagementConfigurationCategoryPlatforms() []string {
	return []string{
		string(DeviceManagementConfigurationCategoryPlatforms_Android),
		string(DeviceManagementConfigurationCategoryPlatforms_IOS),
		string(DeviceManagementConfigurationCategoryPlatforms_Linux),
		string(DeviceManagementConfigurationCategoryPlatforms_MacOS),
		string(DeviceManagementConfigurationCategoryPlatforms_None),
		string(DeviceManagementConfigurationCategoryPlatforms_Windows10),
		string(DeviceManagementConfigurationCategoryPlatforms_Windows10X),
	}
}

func (s *DeviceManagementConfigurationCategoryPlatforms) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationCategoryPlatforms(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationCategoryPlatforms(input string) (*DeviceManagementConfigurationCategoryPlatforms, error) {
	vals := map[string]DeviceManagementConfigurationCategoryPlatforms{
		"android":    DeviceManagementConfigurationCategoryPlatforms_Android,
		"ios":        DeviceManagementConfigurationCategoryPlatforms_IOS,
		"linux":      DeviceManagementConfigurationCategoryPlatforms_Linux,
		"macos":      DeviceManagementConfigurationCategoryPlatforms_MacOS,
		"none":       DeviceManagementConfigurationCategoryPlatforms_None,
		"windows10":  DeviceManagementConfigurationCategoryPlatforms_Windows10,
		"windows10x": DeviceManagementConfigurationCategoryPlatforms_Windows10X,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationCategoryPlatforms(input)
	return &out, nil
}
