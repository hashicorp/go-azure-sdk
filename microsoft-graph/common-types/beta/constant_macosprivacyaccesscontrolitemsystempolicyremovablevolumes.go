package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSPrivacyAccessControlItemSystemPolicyRemovableVolumes string

const (
	MacOSPrivacyAccessControlItemSystemPolicyRemovableVolumes_Disabled      MacOSPrivacyAccessControlItemSystemPolicyRemovableVolumes = "disabled"
	MacOSPrivacyAccessControlItemSystemPolicyRemovableVolumes_Enabled       MacOSPrivacyAccessControlItemSystemPolicyRemovableVolumes = "enabled"
	MacOSPrivacyAccessControlItemSystemPolicyRemovableVolumes_NotConfigured MacOSPrivacyAccessControlItemSystemPolicyRemovableVolumes = "notConfigured"
)

func PossibleValuesForMacOSPrivacyAccessControlItemSystemPolicyRemovableVolumes() []string {
	return []string{
		string(MacOSPrivacyAccessControlItemSystemPolicyRemovableVolumes_Disabled),
		string(MacOSPrivacyAccessControlItemSystemPolicyRemovableVolumes_Enabled),
		string(MacOSPrivacyAccessControlItemSystemPolicyRemovableVolumes_NotConfigured),
	}
}

func (s *MacOSPrivacyAccessControlItemSystemPolicyRemovableVolumes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSPrivacyAccessControlItemSystemPolicyRemovableVolumes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSPrivacyAccessControlItemSystemPolicyRemovableVolumes(input string) (*MacOSPrivacyAccessControlItemSystemPolicyRemovableVolumes, error) {
	vals := map[string]MacOSPrivacyAccessControlItemSystemPolicyRemovableVolumes{
		"disabled":      MacOSPrivacyAccessControlItemSystemPolicyRemovableVolumes_Disabled,
		"enabled":       MacOSPrivacyAccessControlItemSystemPolicyRemovableVolumes_Enabled,
		"notconfigured": MacOSPrivacyAccessControlItemSystemPolicyRemovableVolumes_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSPrivacyAccessControlItemSystemPolicyRemovableVolumes(input)
	return &out, nil
}
