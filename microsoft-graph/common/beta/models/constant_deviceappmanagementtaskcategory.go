package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceAppManagementTaskCategory string

const (
	DeviceAppManagementTaskCategoryadvancedThreatProtection DeviceAppManagementTaskCategory = "AdvancedThreatProtection"
	DeviceAppManagementTaskCategoryunknown                  DeviceAppManagementTaskCategory = "Unknown"
)

func PossibleValuesForDeviceAppManagementTaskCategory() []string {
	return []string{
		string(DeviceAppManagementTaskCategoryadvancedThreatProtection),
		string(DeviceAppManagementTaskCategoryunknown),
	}
}

func parseDeviceAppManagementTaskCategory(input string) (*DeviceAppManagementTaskCategory, error) {
	vals := map[string]DeviceAppManagementTaskCategory{
		"advancedthreatprotection": DeviceAppManagementTaskCategoryadvancedThreatProtection,
		"unknown":                  DeviceAppManagementTaskCategoryunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceAppManagementTaskCategory(input)
	return &out, nil
}
