package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerGeneralDeviceConfigurationEnrollmentProfile string

const (
	AndroidDeviceOwnerGeneralDeviceConfigurationEnrollmentProfile_DedicatedDevice AndroidDeviceOwnerGeneralDeviceConfigurationEnrollmentProfile = "dedicatedDevice"
	AndroidDeviceOwnerGeneralDeviceConfigurationEnrollmentProfile_FullyManaged    AndroidDeviceOwnerGeneralDeviceConfigurationEnrollmentProfile = "fullyManaged"
	AndroidDeviceOwnerGeneralDeviceConfigurationEnrollmentProfile_NotConfigured   AndroidDeviceOwnerGeneralDeviceConfigurationEnrollmentProfile = "notConfigured"
)

func PossibleValuesForAndroidDeviceOwnerGeneralDeviceConfigurationEnrollmentProfile() []string {
	return []string{
		string(AndroidDeviceOwnerGeneralDeviceConfigurationEnrollmentProfile_DedicatedDevice),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationEnrollmentProfile_FullyManaged),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationEnrollmentProfile_NotConfigured),
	}
}

func (s *AndroidDeviceOwnerGeneralDeviceConfigurationEnrollmentProfile) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerGeneralDeviceConfigurationEnrollmentProfile(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerGeneralDeviceConfigurationEnrollmentProfile(input string) (*AndroidDeviceOwnerGeneralDeviceConfigurationEnrollmentProfile, error) {
	vals := map[string]AndroidDeviceOwnerGeneralDeviceConfigurationEnrollmentProfile{
		"dedicateddevice": AndroidDeviceOwnerGeneralDeviceConfigurationEnrollmentProfile_DedicatedDevice,
		"fullymanaged":    AndroidDeviceOwnerGeneralDeviceConfigurationEnrollmentProfile_FullyManaged,
		"notconfigured":   AndroidDeviceOwnerGeneralDeviceConfigurationEnrollmentProfile_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerGeneralDeviceConfigurationEnrollmentProfile(input)
	return &out, nil
}
