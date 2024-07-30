package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetedManagedAppConfigurationTargetedAppManagementLevels string

const (
	TargetedManagedAppConfigurationTargetedAppManagementLevels_AndroidEnterprise                                      TargetedManagedAppConfigurationTargetedAppManagementLevels = "androidEnterprise"
	TargetedManagedAppConfigurationTargetedAppManagementLevels_AndroidEnterpriseDedicatedDevicesWithAzureAdSharedMode TargetedManagedAppConfigurationTargetedAppManagementLevels = "androidEnterpriseDedicatedDevicesWithAzureAdSharedMode"
	TargetedManagedAppConfigurationTargetedAppManagementLevels_AndroidOpenSourceProjectUserAssociated                 TargetedManagedAppConfigurationTargetedAppManagementLevels = "androidOpenSourceProjectUserAssociated"
	TargetedManagedAppConfigurationTargetedAppManagementLevels_AndroidOpenSourceProjectUserless                       TargetedManagedAppConfigurationTargetedAppManagementLevels = "androidOpenSourceProjectUserless"
	TargetedManagedAppConfigurationTargetedAppManagementLevels_Mdm                                                    TargetedManagedAppConfigurationTargetedAppManagementLevels = "mdm"
	TargetedManagedAppConfigurationTargetedAppManagementLevels_Unmanaged                                              TargetedManagedAppConfigurationTargetedAppManagementLevels = "unmanaged"
	TargetedManagedAppConfigurationTargetedAppManagementLevels_Unspecified                                            TargetedManagedAppConfigurationTargetedAppManagementLevels = "unspecified"
)

func PossibleValuesForTargetedManagedAppConfigurationTargetedAppManagementLevels() []string {
	return []string{
		string(TargetedManagedAppConfigurationTargetedAppManagementLevels_AndroidEnterprise),
		string(TargetedManagedAppConfigurationTargetedAppManagementLevels_AndroidEnterpriseDedicatedDevicesWithAzureAdSharedMode),
		string(TargetedManagedAppConfigurationTargetedAppManagementLevels_AndroidOpenSourceProjectUserAssociated),
		string(TargetedManagedAppConfigurationTargetedAppManagementLevels_AndroidOpenSourceProjectUserless),
		string(TargetedManagedAppConfigurationTargetedAppManagementLevels_Mdm),
		string(TargetedManagedAppConfigurationTargetedAppManagementLevels_Unmanaged),
		string(TargetedManagedAppConfigurationTargetedAppManagementLevels_Unspecified),
	}
}

func (s *TargetedManagedAppConfigurationTargetedAppManagementLevels) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTargetedManagedAppConfigurationTargetedAppManagementLevels(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTargetedManagedAppConfigurationTargetedAppManagementLevels(input string) (*TargetedManagedAppConfigurationTargetedAppManagementLevels, error) {
	vals := map[string]TargetedManagedAppConfigurationTargetedAppManagementLevels{
		"androidenterprise": TargetedManagedAppConfigurationTargetedAppManagementLevels_AndroidEnterprise,
		"androidenterprisededicateddeviceswithazureadsharedmode": TargetedManagedAppConfigurationTargetedAppManagementLevels_AndroidEnterpriseDedicatedDevicesWithAzureAdSharedMode,
		"androidopensourceprojectuserassociated":                 TargetedManagedAppConfigurationTargetedAppManagementLevels_AndroidOpenSourceProjectUserAssociated,
		"androidopensourceprojectuserless":                       TargetedManagedAppConfigurationTargetedAppManagementLevels_AndroidOpenSourceProjectUserless,
		"mdm":                                                    TargetedManagedAppConfigurationTargetedAppManagementLevels_Mdm,
		"unmanaged":                                              TargetedManagedAppConfigurationTargetedAppManagementLevels_Unmanaged,
		"unspecified":                                            TargetedManagedAppConfigurationTargetedAppManagementLevels_Unspecified,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TargetedManagedAppConfigurationTargetedAppManagementLevels(input)
	return &out, nil
}
