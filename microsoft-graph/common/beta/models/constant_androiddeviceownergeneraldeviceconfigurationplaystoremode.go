package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerGeneralDeviceConfigurationPlayStoreMode string

const (
	AndroidDeviceOwnerGeneralDeviceConfigurationPlayStoreModeallowList     AndroidDeviceOwnerGeneralDeviceConfigurationPlayStoreMode = "AllowList"
	AndroidDeviceOwnerGeneralDeviceConfigurationPlayStoreModeblockList     AndroidDeviceOwnerGeneralDeviceConfigurationPlayStoreMode = "BlockList"
	AndroidDeviceOwnerGeneralDeviceConfigurationPlayStoreModenotConfigured AndroidDeviceOwnerGeneralDeviceConfigurationPlayStoreMode = "NotConfigured"
)

func PossibleValuesForAndroidDeviceOwnerGeneralDeviceConfigurationPlayStoreMode() []string {
	return []string{
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPlayStoreModeallowList),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPlayStoreModeblockList),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPlayStoreModenotConfigured),
	}
}

func parseAndroidDeviceOwnerGeneralDeviceConfigurationPlayStoreMode(input string) (*AndroidDeviceOwnerGeneralDeviceConfigurationPlayStoreMode, error) {
	vals := map[string]AndroidDeviceOwnerGeneralDeviceConfigurationPlayStoreMode{
		"allowlist":     AndroidDeviceOwnerGeneralDeviceConfigurationPlayStoreModeallowList,
		"blocklist":     AndroidDeviceOwnerGeneralDeviceConfigurationPlayStoreModeblockList,
		"notconfigured": AndroidDeviceOwnerGeneralDeviceConfigurationPlayStoreModenotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerGeneralDeviceConfigurationPlayStoreMode(input)
	return &out, nil
}
