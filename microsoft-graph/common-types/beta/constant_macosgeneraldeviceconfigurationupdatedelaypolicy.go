package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSGeneralDeviceConfigurationUpdateDelayPolicy string

const (
	MacOSGeneralDeviceConfigurationUpdateDelayPolicy_DelayAppUpdateVisibility     MacOSGeneralDeviceConfigurationUpdateDelayPolicy = "delayAppUpdateVisibility"
	MacOSGeneralDeviceConfigurationUpdateDelayPolicy_DelayMajorOsUpdateVisibility MacOSGeneralDeviceConfigurationUpdateDelayPolicy = "delayMajorOsUpdateVisibility"
	MacOSGeneralDeviceConfigurationUpdateDelayPolicy_DelayOSUpdateVisibility      MacOSGeneralDeviceConfigurationUpdateDelayPolicy = "delayOSUpdateVisibility"
	MacOSGeneralDeviceConfigurationUpdateDelayPolicy_None                         MacOSGeneralDeviceConfigurationUpdateDelayPolicy = "none"
)

func PossibleValuesForMacOSGeneralDeviceConfigurationUpdateDelayPolicy() []string {
	return []string{
		string(MacOSGeneralDeviceConfigurationUpdateDelayPolicy_DelayAppUpdateVisibility),
		string(MacOSGeneralDeviceConfigurationUpdateDelayPolicy_DelayMajorOsUpdateVisibility),
		string(MacOSGeneralDeviceConfigurationUpdateDelayPolicy_DelayOSUpdateVisibility),
		string(MacOSGeneralDeviceConfigurationUpdateDelayPolicy_None),
	}
}

func (s *MacOSGeneralDeviceConfigurationUpdateDelayPolicy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSGeneralDeviceConfigurationUpdateDelayPolicy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSGeneralDeviceConfigurationUpdateDelayPolicy(input string) (*MacOSGeneralDeviceConfigurationUpdateDelayPolicy, error) {
	vals := map[string]MacOSGeneralDeviceConfigurationUpdateDelayPolicy{
		"delayappupdatevisibility":     MacOSGeneralDeviceConfigurationUpdateDelayPolicy_DelayAppUpdateVisibility,
		"delaymajorosupdatevisibility": MacOSGeneralDeviceConfigurationUpdateDelayPolicy_DelayMajorOsUpdateVisibility,
		"delayosupdatevisibility":      MacOSGeneralDeviceConfigurationUpdateDelayPolicy_DelayOSUpdateVisibility,
		"none":                         MacOSGeneralDeviceConfigurationUpdateDelayPolicy_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSGeneralDeviceConfigurationUpdateDelayPolicy(input)
	return &out, nil
}
