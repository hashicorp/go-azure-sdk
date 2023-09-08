package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSPrivacyAccessControlItemSystemPolicyRemovableVolumes string

const (
	MacOSPrivacyAccessControlItemSystemPolicyRemovableVolumesdisabled      MacOSPrivacyAccessControlItemSystemPolicyRemovableVolumes = "Disabled"
	MacOSPrivacyAccessControlItemSystemPolicyRemovableVolumesenabled       MacOSPrivacyAccessControlItemSystemPolicyRemovableVolumes = "Enabled"
	MacOSPrivacyAccessControlItemSystemPolicyRemovableVolumesnotConfigured MacOSPrivacyAccessControlItemSystemPolicyRemovableVolumes = "NotConfigured"
)

func PossibleValuesForMacOSPrivacyAccessControlItemSystemPolicyRemovableVolumes() []string {
	return []string{
		string(MacOSPrivacyAccessControlItemSystemPolicyRemovableVolumesdisabled),
		string(MacOSPrivacyAccessControlItemSystemPolicyRemovableVolumesenabled),
		string(MacOSPrivacyAccessControlItemSystemPolicyRemovableVolumesnotConfigured),
	}
}

func parseMacOSPrivacyAccessControlItemSystemPolicyRemovableVolumes(input string) (*MacOSPrivacyAccessControlItemSystemPolicyRemovableVolumes, error) {
	vals := map[string]MacOSPrivacyAccessControlItemSystemPolicyRemovableVolumes{
		"disabled":      MacOSPrivacyAccessControlItemSystemPolicyRemovableVolumesdisabled,
		"enabled":       MacOSPrivacyAccessControlItemSystemPolicyRemovableVolumesenabled,
		"notconfigured": MacOSPrivacyAccessControlItemSystemPolicyRemovableVolumesnotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSPrivacyAccessControlItemSystemPolicyRemovableVolumes(input)
	return &out, nil
}
