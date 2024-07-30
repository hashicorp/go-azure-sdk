package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionTargetedAppManagementLevels string

const (
	AndroidManagedAppProtectionTargetedAppManagementLevels_AndroidEnterprise                                      AndroidManagedAppProtectionTargetedAppManagementLevels = "androidEnterprise"
	AndroidManagedAppProtectionTargetedAppManagementLevels_AndroidEnterpriseDedicatedDevicesWithAzureAdSharedMode AndroidManagedAppProtectionTargetedAppManagementLevels = "androidEnterpriseDedicatedDevicesWithAzureAdSharedMode"
	AndroidManagedAppProtectionTargetedAppManagementLevels_AndroidOpenSourceProjectUserAssociated                 AndroidManagedAppProtectionTargetedAppManagementLevels = "androidOpenSourceProjectUserAssociated"
	AndroidManagedAppProtectionTargetedAppManagementLevels_AndroidOpenSourceProjectUserless                       AndroidManagedAppProtectionTargetedAppManagementLevels = "androidOpenSourceProjectUserless"
	AndroidManagedAppProtectionTargetedAppManagementLevels_Mdm                                                    AndroidManagedAppProtectionTargetedAppManagementLevels = "mdm"
	AndroidManagedAppProtectionTargetedAppManagementLevels_Unmanaged                                              AndroidManagedAppProtectionTargetedAppManagementLevels = "unmanaged"
	AndroidManagedAppProtectionTargetedAppManagementLevels_Unspecified                                            AndroidManagedAppProtectionTargetedAppManagementLevels = "unspecified"
)

func PossibleValuesForAndroidManagedAppProtectionTargetedAppManagementLevels() []string {
	return []string{
		string(AndroidManagedAppProtectionTargetedAppManagementLevels_AndroidEnterprise),
		string(AndroidManagedAppProtectionTargetedAppManagementLevels_AndroidEnterpriseDedicatedDevicesWithAzureAdSharedMode),
		string(AndroidManagedAppProtectionTargetedAppManagementLevels_AndroidOpenSourceProjectUserAssociated),
		string(AndroidManagedAppProtectionTargetedAppManagementLevels_AndroidOpenSourceProjectUserless),
		string(AndroidManagedAppProtectionTargetedAppManagementLevels_Mdm),
		string(AndroidManagedAppProtectionTargetedAppManagementLevels_Unmanaged),
		string(AndroidManagedAppProtectionTargetedAppManagementLevels_Unspecified),
	}
}

func (s *AndroidManagedAppProtectionTargetedAppManagementLevels) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidManagedAppProtectionTargetedAppManagementLevels(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidManagedAppProtectionTargetedAppManagementLevels(input string) (*AndroidManagedAppProtectionTargetedAppManagementLevels, error) {
	vals := map[string]AndroidManagedAppProtectionTargetedAppManagementLevels{
		"androidenterprise": AndroidManagedAppProtectionTargetedAppManagementLevels_AndroidEnterprise,
		"androidenterprisededicateddeviceswithazureadsharedmode": AndroidManagedAppProtectionTargetedAppManagementLevels_AndroidEnterpriseDedicatedDevicesWithAzureAdSharedMode,
		"androidopensourceprojectuserassociated":                 AndroidManagedAppProtectionTargetedAppManagementLevels_AndroidOpenSourceProjectUserAssociated,
		"androidopensourceprojectuserless":                       AndroidManagedAppProtectionTargetedAppManagementLevels_AndroidOpenSourceProjectUserless,
		"mdm":                                                    AndroidManagedAppProtectionTargetedAppManagementLevels_Mdm,
		"unmanaged":                                              AndroidManagedAppProtectionTargetedAppManagementLevels_Unmanaged,
		"unspecified":                                            AndroidManagedAppProtectionTargetedAppManagementLevels_Unspecified,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionTargetedAppManagementLevels(input)
	return &out, nil
}
