package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSDeviceFeaturesConfigurationContentCachingType string

const (
	MacOSDeviceFeaturesConfigurationContentCachingTypenotConfigured     MacOSDeviceFeaturesConfigurationContentCachingType = "NotConfigured"
	MacOSDeviceFeaturesConfigurationContentCachingTypesharedContentOnly MacOSDeviceFeaturesConfigurationContentCachingType = "SharedContentOnly"
	MacOSDeviceFeaturesConfigurationContentCachingTypeuserContentOnly   MacOSDeviceFeaturesConfigurationContentCachingType = "UserContentOnly"
)

func PossibleValuesForMacOSDeviceFeaturesConfigurationContentCachingType() []string {
	return []string{
		string(MacOSDeviceFeaturesConfigurationContentCachingTypenotConfigured),
		string(MacOSDeviceFeaturesConfigurationContentCachingTypesharedContentOnly),
		string(MacOSDeviceFeaturesConfigurationContentCachingTypeuserContentOnly),
	}
}

func parseMacOSDeviceFeaturesConfigurationContentCachingType(input string) (*MacOSDeviceFeaturesConfigurationContentCachingType, error) {
	vals := map[string]MacOSDeviceFeaturesConfigurationContentCachingType{
		"notconfigured":     MacOSDeviceFeaturesConfigurationContentCachingTypenotConfigured,
		"sharedcontentonly": MacOSDeviceFeaturesConfigurationContentCachingTypesharedContentOnly,
		"usercontentonly":   MacOSDeviceFeaturesConfigurationContentCachingTypeuserContentOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSDeviceFeaturesConfigurationContentCachingType(input)
	return &out, nil
}
