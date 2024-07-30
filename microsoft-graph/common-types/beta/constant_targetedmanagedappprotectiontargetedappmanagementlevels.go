package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetedManagedAppProtectionTargetedAppManagementLevels string

const (
	TargetedManagedAppProtectionTargetedAppManagementLevels_AndroidEnterprise                                      TargetedManagedAppProtectionTargetedAppManagementLevels = "androidEnterprise"
	TargetedManagedAppProtectionTargetedAppManagementLevels_AndroidEnterpriseDedicatedDevicesWithAzureAdSharedMode TargetedManagedAppProtectionTargetedAppManagementLevels = "androidEnterpriseDedicatedDevicesWithAzureAdSharedMode"
	TargetedManagedAppProtectionTargetedAppManagementLevels_AndroidOpenSourceProjectUserAssociated                 TargetedManagedAppProtectionTargetedAppManagementLevels = "androidOpenSourceProjectUserAssociated"
	TargetedManagedAppProtectionTargetedAppManagementLevels_AndroidOpenSourceProjectUserless                       TargetedManagedAppProtectionTargetedAppManagementLevels = "androidOpenSourceProjectUserless"
	TargetedManagedAppProtectionTargetedAppManagementLevels_Mdm                                                    TargetedManagedAppProtectionTargetedAppManagementLevels = "mdm"
	TargetedManagedAppProtectionTargetedAppManagementLevels_Unmanaged                                              TargetedManagedAppProtectionTargetedAppManagementLevels = "unmanaged"
	TargetedManagedAppProtectionTargetedAppManagementLevels_Unspecified                                            TargetedManagedAppProtectionTargetedAppManagementLevels = "unspecified"
)

func PossibleValuesForTargetedManagedAppProtectionTargetedAppManagementLevels() []string {
	return []string{
		string(TargetedManagedAppProtectionTargetedAppManagementLevels_AndroidEnterprise),
		string(TargetedManagedAppProtectionTargetedAppManagementLevels_AndroidEnterpriseDedicatedDevicesWithAzureAdSharedMode),
		string(TargetedManagedAppProtectionTargetedAppManagementLevels_AndroidOpenSourceProjectUserAssociated),
		string(TargetedManagedAppProtectionTargetedAppManagementLevels_AndroidOpenSourceProjectUserless),
		string(TargetedManagedAppProtectionTargetedAppManagementLevels_Mdm),
		string(TargetedManagedAppProtectionTargetedAppManagementLevels_Unmanaged),
		string(TargetedManagedAppProtectionTargetedAppManagementLevels_Unspecified),
	}
}

func (s *TargetedManagedAppProtectionTargetedAppManagementLevels) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTargetedManagedAppProtectionTargetedAppManagementLevels(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTargetedManagedAppProtectionTargetedAppManagementLevels(input string) (*TargetedManagedAppProtectionTargetedAppManagementLevels, error) {
	vals := map[string]TargetedManagedAppProtectionTargetedAppManagementLevels{
		"androidenterprise": TargetedManagedAppProtectionTargetedAppManagementLevels_AndroidEnterprise,
		"androidenterprisededicateddeviceswithazureadsharedmode": TargetedManagedAppProtectionTargetedAppManagementLevels_AndroidEnterpriseDedicatedDevicesWithAzureAdSharedMode,
		"androidopensourceprojectuserassociated":                 TargetedManagedAppProtectionTargetedAppManagementLevels_AndroidOpenSourceProjectUserAssociated,
		"androidopensourceprojectuserless":                       TargetedManagedAppProtectionTargetedAppManagementLevels_AndroidOpenSourceProjectUserless,
		"mdm":                                                    TargetedManagedAppProtectionTargetedAppManagementLevels_Mdm,
		"unmanaged":                                              TargetedManagedAppProtectionTargetedAppManagementLevels_Unmanaged,
		"unspecified":                                            TargetedManagedAppProtectionTargetedAppManagementLevels_Unspecified,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TargetedManagedAppProtectionTargetedAppManagementLevels(input)
	return &out, nil
}
