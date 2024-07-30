package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerGeneralDeviceConfigurationCrossProfilePoliciesAllowDataSharing string

const (
	AndroidDeviceOwnerGeneralDeviceConfigurationCrossProfilePoliciesAllowDataSharing_CrossProfileDataSharingAllowed       AndroidDeviceOwnerGeneralDeviceConfigurationCrossProfilePoliciesAllowDataSharing = "crossProfileDataSharingAllowed"
	AndroidDeviceOwnerGeneralDeviceConfigurationCrossProfilePoliciesAllowDataSharing_CrossProfileDataSharingBlocked       AndroidDeviceOwnerGeneralDeviceConfigurationCrossProfilePoliciesAllowDataSharing = "crossProfileDataSharingBlocked"
	AndroidDeviceOwnerGeneralDeviceConfigurationCrossProfilePoliciesAllowDataSharing_DataSharingFromWorkToPersonalBlocked AndroidDeviceOwnerGeneralDeviceConfigurationCrossProfilePoliciesAllowDataSharing = "dataSharingFromWorkToPersonalBlocked"
	AndroidDeviceOwnerGeneralDeviceConfigurationCrossProfilePoliciesAllowDataSharing_NotConfigured                        AndroidDeviceOwnerGeneralDeviceConfigurationCrossProfilePoliciesAllowDataSharing = "notConfigured"
	AndroidDeviceOwnerGeneralDeviceConfigurationCrossProfilePoliciesAllowDataSharing_UnkownFutureValue                    AndroidDeviceOwnerGeneralDeviceConfigurationCrossProfilePoliciesAllowDataSharing = "unkownFutureValue"
)

func PossibleValuesForAndroidDeviceOwnerGeneralDeviceConfigurationCrossProfilePoliciesAllowDataSharing() []string {
	return []string{
		string(AndroidDeviceOwnerGeneralDeviceConfigurationCrossProfilePoliciesAllowDataSharing_CrossProfileDataSharingAllowed),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationCrossProfilePoliciesAllowDataSharing_CrossProfileDataSharingBlocked),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationCrossProfilePoliciesAllowDataSharing_DataSharingFromWorkToPersonalBlocked),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationCrossProfilePoliciesAllowDataSharing_NotConfigured),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationCrossProfilePoliciesAllowDataSharing_UnkownFutureValue),
	}
}

func (s *AndroidDeviceOwnerGeneralDeviceConfigurationCrossProfilePoliciesAllowDataSharing) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerGeneralDeviceConfigurationCrossProfilePoliciesAllowDataSharing(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerGeneralDeviceConfigurationCrossProfilePoliciesAllowDataSharing(input string) (*AndroidDeviceOwnerGeneralDeviceConfigurationCrossProfilePoliciesAllowDataSharing, error) {
	vals := map[string]AndroidDeviceOwnerGeneralDeviceConfigurationCrossProfilePoliciesAllowDataSharing{
		"crossprofiledatasharingallowed":       AndroidDeviceOwnerGeneralDeviceConfigurationCrossProfilePoliciesAllowDataSharing_CrossProfileDataSharingAllowed,
		"crossprofiledatasharingblocked":       AndroidDeviceOwnerGeneralDeviceConfigurationCrossProfilePoliciesAllowDataSharing_CrossProfileDataSharingBlocked,
		"datasharingfromworktopersonalblocked": AndroidDeviceOwnerGeneralDeviceConfigurationCrossProfilePoliciesAllowDataSharing_DataSharingFromWorkToPersonalBlocked,
		"notconfigured":                        AndroidDeviceOwnerGeneralDeviceConfigurationCrossProfilePoliciesAllowDataSharing_NotConfigured,
		"unkownfuturevalue":                    AndroidDeviceOwnerGeneralDeviceConfigurationCrossProfilePoliciesAllowDataSharing_UnkownFutureValue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerGeneralDeviceConfigurationCrossProfilePoliciesAllowDataSharing(input)
	return &out, nil
}
