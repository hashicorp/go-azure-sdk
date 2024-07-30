package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosGeneralDeviceConfigurationAppsVisibilityListType string

const (
	IosGeneralDeviceConfigurationAppsVisibilityListType_AppsInListCompliant    IosGeneralDeviceConfigurationAppsVisibilityListType = "appsInListCompliant"
	IosGeneralDeviceConfigurationAppsVisibilityListType_AppsNotInListCompliant IosGeneralDeviceConfigurationAppsVisibilityListType = "appsNotInListCompliant"
	IosGeneralDeviceConfigurationAppsVisibilityListType_None                   IosGeneralDeviceConfigurationAppsVisibilityListType = "none"
)

func PossibleValuesForIosGeneralDeviceConfigurationAppsVisibilityListType() []string {
	return []string{
		string(IosGeneralDeviceConfigurationAppsVisibilityListType_AppsInListCompliant),
		string(IosGeneralDeviceConfigurationAppsVisibilityListType_AppsNotInListCompliant),
		string(IosGeneralDeviceConfigurationAppsVisibilityListType_None),
	}
}

func (s *IosGeneralDeviceConfigurationAppsVisibilityListType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosGeneralDeviceConfigurationAppsVisibilityListType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosGeneralDeviceConfigurationAppsVisibilityListType(input string) (*IosGeneralDeviceConfigurationAppsVisibilityListType, error) {
	vals := map[string]IosGeneralDeviceConfigurationAppsVisibilityListType{
		"appsinlistcompliant":    IosGeneralDeviceConfigurationAppsVisibilityListType_AppsInListCompliant,
		"appsnotinlistcompliant": IosGeneralDeviceConfigurationAppsVisibilityListType_AppsNotInListCompliant,
		"none":                   IosGeneralDeviceConfigurationAppsVisibilityListType_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosGeneralDeviceConfigurationAppsVisibilityListType(input)
	return &out, nil
}
