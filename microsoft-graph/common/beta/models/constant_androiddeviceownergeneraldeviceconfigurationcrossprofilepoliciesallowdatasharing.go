package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerGeneralDeviceConfigurationCrossProfilePoliciesAllowDataSharing string

const (
	AndroidDeviceOwnerGeneralDeviceConfigurationCrossProfilePoliciesAllowDataSharingcrossProfileDataSharingAllowed       AndroidDeviceOwnerGeneralDeviceConfigurationCrossProfilePoliciesAllowDataSharing = "CrossProfileDataSharingAllowed"
	AndroidDeviceOwnerGeneralDeviceConfigurationCrossProfilePoliciesAllowDataSharingcrossProfileDataSharingBlocked       AndroidDeviceOwnerGeneralDeviceConfigurationCrossProfilePoliciesAllowDataSharing = "CrossProfileDataSharingBlocked"
	AndroidDeviceOwnerGeneralDeviceConfigurationCrossProfilePoliciesAllowDataSharingdataSharingFromWorkToPersonalBlocked AndroidDeviceOwnerGeneralDeviceConfigurationCrossProfilePoliciesAllowDataSharing = "DataSharingFromWorkToPersonalBlocked"
	AndroidDeviceOwnerGeneralDeviceConfigurationCrossProfilePoliciesAllowDataSharingnotConfigured                        AndroidDeviceOwnerGeneralDeviceConfigurationCrossProfilePoliciesAllowDataSharing = "NotConfigured"
	AndroidDeviceOwnerGeneralDeviceConfigurationCrossProfilePoliciesAllowDataSharingunkownFutureValue                    AndroidDeviceOwnerGeneralDeviceConfigurationCrossProfilePoliciesAllowDataSharing = "UnkownFutureValue"
)

func PossibleValuesForAndroidDeviceOwnerGeneralDeviceConfigurationCrossProfilePoliciesAllowDataSharing() []string {
	return []string{
		string(AndroidDeviceOwnerGeneralDeviceConfigurationCrossProfilePoliciesAllowDataSharingcrossProfileDataSharingAllowed),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationCrossProfilePoliciesAllowDataSharingcrossProfileDataSharingBlocked),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationCrossProfilePoliciesAllowDataSharingdataSharingFromWorkToPersonalBlocked),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationCrossProfilePoliciesAllowDataSharingnotConfigured),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationCrossProfilePoliciesAllowDataSharingunkownFutureValue),
	}
}

func parseAndroidDeviceOwnerGeneralDeviceConfigurationCrossProfilePoliciesAllowDataSharing(input string) (*AndroidDeviceOwnerGeneralDeviceConfigurationCrossProfilePoliciesAllowDataSharing, error) {
	vals := map[string]AndroidDeviceOwnerGeneralDeviceConfigurationCrossProfilePoliciesAllowDataSharing{
		"crossprofiledatasharingallowed":       AndroidDeviceOwnerGeneralDeviceConfigurationCrossProfilePoliciesAllowDataSharingcrossProfileDataSharingAllowed,
		"crossprofiledatasharingblocked":       AndroidDeviceOwnerGeneralDeviceConfigurationCrossProfilePoliciesAllowDataSharingcrossProfileDataSharingBlocked,
		"datasharingfromworktopersonalblocked": AndroidDeviceOwnerGeneralDeviceConfigurationCrossProfilePoliciesAllowDataSharingdataSharingFromWorkToPersonalBlocked,
		"notconfigured":                        AndroidDeviceOwnerGeneralDeviceConfigurationCrossProfilePoliciesAllowDataSharingnotConfigured,
		"unkownfuturevalue":                    AndroidDeviceOwnerGeneralDeviceConfigurationCrossProfilePoliciesAllowDataSharingunkownFutureValue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerGeneralDeviceConfigurationCrossProfilePoliciesAllowDataSharing(input)
	return &out, nil
}
