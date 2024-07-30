package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSPrivacyAccessControlItemSystemPolicyNetworkVolumes string

const (
	MacOSPrivacyAccessControlItemSystemPolicyNetworkVolumes_Disabled      MacOSPrivacyAccessControlItemSystemPolicyNetworkVolumes = "disabled"
	MacOSPrivacyAccessControlItemSystemPolicyNetworkVolumes_Enabled       MacOSPrivacyAccessControlItemSystemPolicyNetworkVolumes = "enabled"
	MacOSPrivacyAccessControlItemSystemPolicyNetworkVolumes_NotConfigured MacOSPrivacyAccessControlItemSystemPolicyNetworkVolumes = "notConfigured"
)

func PossibleValuesForMacOSPrivacyAccessControlItemSystemPolicyNetworkVolumes() []string {
	return []string{
		string(MacOSPrivacyAccessControlItemSystemPolicyNetworkVolumes_Disabled),
		string(MacOSPrivacyAccessControlItemSystemPolicyNetworkVolumes_Enabled),
		string(MacOSPrivacyAccessControlItemSystemPolicyNetworkVolumes_NotConfigured),
	}
}

func (s *MacOSPrivacyAccessControlItemSystemPolicyNetworkVolumes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSPrivacyAccessControlItemSystemPolicyNetworkVolumes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSPrivacyAccessControlItemSystemPolicyNetworkVolumes(input string) (*MacOSPrivacyAccessControlItemSystemPolicyNetworkVolumes, error) {
	vals := map[string]MacOSPrivacyAccessControlItemSystemPolicyNetworkVolumes{
		"disabled":      MacOSPrivacyAccessControlItemSystemPolicyNetworkVolumes_Disabled,
		"enabled":       MacOSPrivacyAccessControlItemSystemPolicyNetworkVolumes_Enabled,
		"notconfigured": MacOSPrivacyAccessControlItemSystemPolicyNetworkVolumes_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSPrivacyAccessControlItemSystemPolicyNetworkVolumes(input)
	return &out, nil
}
