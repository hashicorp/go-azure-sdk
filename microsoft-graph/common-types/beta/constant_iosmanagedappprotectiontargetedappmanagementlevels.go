package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosManagedAppProtectionTargetedAppManagementLevels string

const (
	IosManagedAppProtectionTargetedAppManagementLevels_AndroidEnterprise                                      IosManagedAppProtectionTargetedAppManagementLevels = "androidEnterprise"
	IosManagedAppProtectionTargetedAppManagementLevels_AndroidEnterpriseDedicatedDevicesWithAzureAdSharedMode IosManagedAppProtectionTargetedAppManagementLevels = "androidEnterpriseDedicatedDevicesWithAzureAdSharedMode"
	IosManagedAppProtectionTargetedAppManagementLevels_AndroidOpenSourceProjectUserAssociated                 IosManagedAppProtectionTargetedAppManagementLevels = "androidOpenSourceProjectUserAssociated"
	IosManagedAppProtectionTargetedAppManagementLevels_AndroidOpenSourceProjectUserless                       IosManagedAppProtectionTargetedAppManagementLevels = "androidOpenSourceProjectUserless"
	IosManagedAppProtectionTargetedAppManagementLevels_Mdm                                                    IosManagedAppProtectionTargetedAppManagementLevels = "mdm"
	IosManagedAppProtectionTargetedAppManagementLevels_Unmanaged                                              IosManagedAppProtectionTargetedAppManagementLevels = "unmanaged"
	IosManagedAppProtectionTargetedAppManagementLevels_Unspecified                                            IosManagedAppProtectionTargetedAppManagementLevels = "unspecified"
)

func PossibleValuesForIosManagedAppProtectionTargetedAppManagementLevels() []string {
	return []string{
		string(IosManagedAppProtectionTargetedAppManagementLevels_AndroidEnterprise),
		string(IosManagedAppProtectionTargetedAppManagementLevels_AndroidEnterpriseDedicatedDevicesWithAzureAdSharedMode),
		string(IosManagedAppProtectionTargetedAppManagementLevels_AndroidOpenSourceProjectUserAssociated),
		string(IosManagedAppProtectionTargetedAppManagementLevels_AndroidOpenSourceProjectUserless),
		string(IosManagedAppProtectionTargetedAppManagementLevels_Mdm),
		string(IosManagedAppProtectionTargetedAppManagementLevels_Unmanaged),
		string(IosManagedAppProtectionTargetedAppManagementLevels_Unspecified),
	}
}

func (s *IosManagedAppProtectionTargetedAppManagementLevels) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosManagedAppProtectionTargetedAppManagementLevels(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosManagedAppProtectionTargetedAppManagementLevels(input string) (*IosManagedAppProtectionTargetedAppManagementLevels, error) {
	vals := map[string]IosManagedAppProtectionTargetedAppManagementLevels{
		"androidenterprise": IosManagedAppProtectionTargetedAppManagementLevels_AndroidEnterprise,
		"androidenterprisededicateddeviceswithazureadsharedmode": IosManagedAppProtectionTargetedAppManagementLevels_AndroidEnterpriseDedicatedDevicesWithAzureAdSharedMode,
		"androidopensourceprojectuserassociated":                 IosManagedAppProtectionTargetedAppManagementLevels_AndroidOpenSourceProjectUserAssociated,
		"androidopensourceprojectuserless":                       IosManagedAppProtectionTargetedAppManagementLevels_AndroidOpenSourceProjectUserless,
		"mdm":                                                    IosManagedAppProtectionTargetedAppManagementLevels_Mdm,
		"unmanaged":                                              IosManagedAppProtectionTargetedAppManagementLevels_Unmanaged,
		"unspecified":                                            IosManagedAppProtectionTargetedAppManagementLevels_Unspecified,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosManagedAppProtectionTargetedAppManagementLevels(input)
	return &out, nil
}
