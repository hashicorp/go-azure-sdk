package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosGeneralDeviceConfigurationKioskModeAppType string

const (
	IosGeneralDeviceConfigurationKioskModeAppType_AppStoreApp   IosGeneralDeviceConfigurationKioskModeAppType = "appStoreApp"
	IosGeneralDeviceConfigurationKioskModeAppType_BuiltInApp    IosGeneralDeviceConfigurationKioskModeAppType = "builtInApp"
	IosGeneralDeviceConfigurationKioskModeAppType_ManagedApp    IosGeneralDeviceConfigurationKioskModeAppType = "managedApp"
	IosGeneralDeviceConfigurationKioskModeAppType_NotConfigured IosGeneralDeviceConfigurationKioskModeAppType = "notConfigured"
)

func PossibleValuesForIosGeneralDeviceConfigurationKioskModeAppType() []string {
	return []string{
		string(IosGeneralDeviceConfigurationKioskModeAppType_AppStoreApp),
		string(IosGeneralDeviceConfigurationKioskModeAppType_BuiltInApp),
		string(IosGeneralDeviceConfigurationKioskModeAppType_ManagedApp),
		string(IosGeneralDeviceConfigurationKioskModeAppType_NotConfigured),
	}
}

func (s *IosGeneralDeviceConfigurationKioskModeAppType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosGeneralDeviceConfigurationKioskModeAppType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosGeneralDeviceConfigurationKioskModeAppType(input string) (*IosGeneralDeviceConfigurationKioskModeAppType, error) {
	vals := map[string]IosGeneralDeviceConfigurationKioskModeAppType{
		"appstoreapp":   IosGeneralDeviceConfigurationKioskModeAppType_AppStoreApp,
		"builtinapp":    IosGeneralDeviceConfigurationKioskModeAppType_BuiltInApp,
		"managedapp":    IosGeneralDeviceConfigurationKioskModeAppType_ManagedApp,
		"notconfigured": IosGeneralDeviceConfigurationKioskModeAppType_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosGeneralDeviceConfigurationKioskModeAppType(input)
	return &out, nil
}
