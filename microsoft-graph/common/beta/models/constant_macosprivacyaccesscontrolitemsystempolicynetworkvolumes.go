package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSPrivacyAccessControlItemSystemPolicyNetworkVolumes string

const (
	MacOSPrivacyAccessControlItemSystemPolicyNetworkVolumesdisabled      MacOSPrivacyAccessControlItemSystemPolicyNetworkVolumes = "Disabled"
	MacOSPrivacyAccessControlItemSystemPolicyNetworkVolumesenabled       MacOSPrivacyAccessControlItemSystemPolicyNetworkVolumes = "Enabled"
	MacOSPrivacyAccessControlItemSystemPolicyNetworkVolumesnotConfigured MacOSPrivacyAccessControlItemSystemPolicyNetworkVolumes = "NotConfigured"
)

func PossibleValuesForMacOSPrivacyAccessControlItemSystemPolicyNetworkVolumes() []string {
	return []string{
		string(MacOSPrivacyAccessControlItemSystemPolicyNetworkVolumesdisabled),
		string(MacOSPrivacyAccessControlItemSystemPolicyNetworkVolumesenabled),
		string(MacOSPrivacyAccessControlItemSystemPolicyNetworkVolumesnotConfigured),
	}
}

func parseMacOSPrivacyAccessControlItemSystemPolicyNetworkVolumes(input string) (*MacOSPrivacyAccessControlItemSystemPolicyNetworkVolumes, error) {
	vals := map[string]MacOSPrivacyAccessControlItemSystemPolicyNetworkVolumes{
		"disabled":      MacOSPrivacyAccessControlItemSystemPolicyNetworkVolumesdisabled,
		"enabled":       MacOSPrivacyAccessControlItemSystemPolicyNetworkVolumesenabled,
		"notconfigured": MacOSPrivacyAccessControlItemSystemPolicyNetworkVolumesnotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSPrivacyAccessControlItemSystemPolicyNetworkVolumes(input)
	return &out, nil
}
