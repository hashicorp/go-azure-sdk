package resourceaccessprofile

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType string

const (
	DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeExclude DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType = "exclude"
	DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeInclude DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType = "include"
	DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeNone    DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType = "none"
)

func PossibleValuesForDeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType() []string {
	return []string{
		string(DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeExclude),
		string(DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeInclude),
		string(DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeNone),
	}
}

func (s *DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType(input string) (*DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType, error) {
	vals := map[string]DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType{
		"exclude": DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeExclude,
		"include": DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeInclude,
		"none":    DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType(input)
	return &out, nil
}

type DeviceManagementResourceAccessProfileAssignmentIntent string

const (
	DeviceManagementResourceAccessProfileAssignmentIntentApply  DeviceManagementResourceAccessProfileAssignmentIntent = "apply"
	DeviceManagementResourceAccessProfileAssignmentIntentRemove DeviceManagementResourceAccessProfileAssignmentIntent = "remove"
)

func PossibleValuesForDeviceManagementResourceAccessProfileAssignmentIntent() []string {
	return []string{
		string(DeviceManagementResourceAccessProfileAssignmentIntentApply),
		string(DeviceManagementResourceAccessProfileAssignmentIntentRemove),
	}
}

func (s *DeviceManagementResourceAccessProfileAssignmentIntent) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementResourceAccessProfileAssignmentIntent(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementResourceAccessProfileAssignmentIntent(input string) (*DeviceManagementResourceAccessProfileAssignmentIntent, error) {
	vals := map[string]DeviceManagementResourceAccessProfileAssignmentIntent{
		"apply":  DeviceManagementResourceAccessProfileAssignmentIntentApply,
		"remove": DeviceManagementResourceAccessProfileAssignmentIntentRemove,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementResourceAccessProfileAssignmentIntent(input)
	return &out, nil
}

type ListResourceAccessProfileQueryByPlatformTypesRequestPlatformType string

const (
	ListResourceAccessProfileQueryByPlatformTypesRequestPlatformTypeAll                ListResourceAccessProfileQueryByPlatformTypesRequestPlatformType = "all"
	ListResourceAccessProfileQueryByPlatformTypesRequestPlatformTypeAndroid            ListResourceAccessProfileQueryByPlatformTypesRequestPlatformType = "android"
	ListResourceAccessProfileQueryByPlatformTypesRequestPlatformTypeAndroidAOSP        ListResourceAccessProfileQueryByPlatformTypesRequestPlatformType = "androidAOSP"
	ListResourceAccessProfileQueryByPlatformTypesRequestPlatformTypeAndroidForWork     ListResourceAccessProfileQueryByPlatformTypesRequestPlatformType = "androidForWork"
	ListResourceAccessProfileQueryByPlatformTypesRequestPlatformTypeAndroidWorkProfile ListResourceAccessProfileQueryByPlatformTypesRequestPlatformType = "androidWorkProfile"
	ListResourceAccessProfileQueryByPlatformTypesRequestPlatformTypeIOS                ListResourceAccessProfileQueryByPlatformTypesRequestPlatformType = "iOS"
	ListResourceAccessProfileQueryByPlatformTypesRequestPlatformTypeMacOS              ListResourceAccessProfileQueryByPlatformTypesRequestPlatformType = "macOS"
	ListResourceAccessProfileQueryByPlatformTypesRequestPlatformTypeWindows10AndLater  ListResourceAccessProfileQueryByPlatformTypesRequestPlatformType = "windows10AndLater"
	ListResourceAccessProfileQueryByPlatformTypesRequestPlatformTypeWindows10XProfile  ListResourceAccessProfileQueryByPlatformTypesRequestPlatformType = "windows10XProfile"
	ListResourceAccessProfileQueryByPlatformTypesRequestPlatformTypeWindows81AndLater  ListResourceAccessProfileQueryByPlatformTypesRequestPlatformType = "windows81AndLater"
	ListResourceAccessProfileQueryByPlatformTypesRequestPlatformTypeWindowsPhone81     ListResourceAccessProfileQueryByPlatformTypesRequestPlatformType = "windowsPhone81"
)

func PossibleValuesForListResourceAccessProfileQueryByPlatformTypesRequestPlatformType() []string {
	return []string{
		string(ListResourceAccessProfileQueryByPlatformTypesRequestPlatformTypeAll),
		string(ListResourceAccessProfileQueryByPlatformTypesRequestPlatformTypeAndroid),
		string(ListResourceAccessProfileQueryByPlatformTypesRequestPlatformTypeAndroidAOSP),
		string(ListResourceAccessProfileQueryByPlatformTypesRequestPlatformTypeAndroidForWork),
		string(ListResourceAccessProfileQueryByPlatformTypesRequestPlatformTypeAndroidWorkProfile),
		string(ListResourceAccessProfileQueryByPlatformTypesRequestPlatformTypeIOS),
		string(ListResourceAccessProfileQueryByPlatformTypesRequestPlatformTypeMacOS),
		string(ListResourceAccessProfileQueryByPlatformTypesRequestPlatformTypeWindows10AndLater),
		string(ListResourceAccessProfileQueryByPlatformTypesRequestPlatformTypeWindows10XProfile),
		string(ListResourceAccessProfileQueryByPlatformTypesRequestPlatformTypeWindows81AndLater),
		string(ListResourceAccessProfileQueryByPlatformTypesRequestPlatformTypeWindowsPhone81),
	}
}

func (s *ListResourceAccessProfileQueryByPlatformTypesRequestPlatformType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseListResourceAccessProfileQueryByPlatformTypesRequestPlatformType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseListResourceAccessProfileQueryByPlatformTypesRequestPlatformType(input string) (*ListResourceAccessProfileQueryByPlatformTypesRequestPlatformType, error) {
	vals := map[string]ListResourceAccessProfileQueryByPlatformTypesRequestPlatformType{
		"all":                ListResourceAccessProfileQueryByPlatformTypesRequestPlatformTypeAll,
		"android":            ListResourceAccessProfileQueryByPlatformTypesRequestPlatformTypeAndroid,
		"androidaosp":        ListResourceAccessProfileQueryByPlatformTypesRequestPlatformTypeAndroidAOSP,
		"androidforwork":     ListResourceAccessProfileQueryByPlatformTypesRequestPlatformTypeAndroidForWork,
		"androidworkprofile": ListResourceAccessProfileQueryByPlatformTypesRequestPlatformTypeAndroidWorkProfile,
		"ios":                ListResourceAccessProfileQueryByPlatformTypesRequestPlatformTypeIOS,
		"macos":              ListResourceAccessProfileQueryByPlatformTypesRequestPlatformTypeMacOS,
		"windows10andlater":  ListResourceAccessProfileQueryByPlatformTypesRequestPlatformTypeWindows10AndLater,
		"windows10xprofile":  ListResourceAccessProfileQueryByPlatformTypesRequestPlatformTypeWindows10XProfile,
		"windows81andlater":  ListResourceAccessProfileQueryByPlatformTypesRequestPlatformTypeWindows81AndLater,
		"windowsphone81":     ListResourceAccessProfileQueryByPlatformTypesRequestPlatformTypeWindowsPhone81,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ListResourceAccessProfileQueryByPlatformTypesRequestPlatformType(input)
	return &out, nil
}
