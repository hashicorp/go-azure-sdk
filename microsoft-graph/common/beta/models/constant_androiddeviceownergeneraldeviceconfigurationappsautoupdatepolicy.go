package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerGeneralDeviceConfigurationAppsAutoUpdatePolicy string

const (
	AndroidDeviceOwnerGeneralDeviceConfigurationAppsAutoUpdatePolicyalways        AndroidDeviceOwnerGeneralDeviceConfigurationAppsAutoUpdatePolicy = "Always"
	AndroidDeviceOwnerGeneralDeviceConfigurationAppsAutoUpdatePolicynever         AndroidDeviceOwnerGeneralDeviceConfigurationAppsAutoUpdatePolicy = "Never"
	AndroidDeviceOwnerGeneralDeviceConfigurationAppsAutoUpdatePolicynotConfigured AndroidDeviceOwnerGeneralDeviceConfigurationAppsAutoUpdatePolicy = "NotConfigured"
	AndroidDeviceOwnerGeneralDeviceConfigurationAppsAutoUpdatePolicyuserChoice    AndroidDeviceOwnerGeneralDeviceConfigurationAppsAutoUpdatePolicy = "UserChoice"
	AndroidDeviceOwnerGeneralDeviceConfigurationAppsAutoUpdatePolicywiFiOnly      AndroidDeviceOwnerGeneralDeviceConfigurationAppsAutoUpdatePolicy = "WiFiOnly"
)

func PossibleValuesForAndroidDeviceOwnerGeneralDeviceConfigurationAppsAutoUpdatePolicy() []string {
	return []string{
		string(AndroidDeviceOwnerGeneralDeviceConfigurationAppsAutoUpdatePolicyalways),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationAppsAutoUpdatePolicynever),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationAppsAutoUpdatePolicynotConfigured),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationAppsAutoUpdatePolicyuserChoice),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationAppsAutoUpdatePolicywiFiOnly),
	}
}

func parseAndroidDeviceOwnerGeneralDeviceConfigurationAppsAutoUpdatePolicy(input string) (*AndroidDeviceOwnerGeneralDeviceConfigurationAppsAutoUpdatePolicy, error) {
	vals := map[string]AndroidDeviceOwnerGeneralDeviceConfigurationAppsAutoUpdatePolicy{
		"always":        AndroidDeviceOwnerGeneralDeviceConfigurationAppsAutoUpdatePolicyalways,
		"never":         AndroidDeviceOwnerGeneralDeviceConfigurationAppsAutoUpdatePolicynever,
		"notconfigured": AndroidDeviceOwnerGeneralDeviceConfigurationAppsAutoUpdatePolicynotConfigured,
		"userchoice":    AndroidDeviceOwnerGeneralDeviceConfigurationAppsAutoUpdatePolicyuserChoice,
		"wifionly":      AndroidDeviceOwnerGeneralDeviceConfigurationAppsAutoUpdatePolicywiFiOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerGeneralDeviceConfigurationAppsAutoUpdatePolicy(input)
	return &out, nil
}
