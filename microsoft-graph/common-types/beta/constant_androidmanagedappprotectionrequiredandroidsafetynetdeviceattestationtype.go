package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType string

const (
	AndroidManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType_BasicIntegrity                       AndroidManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType = "basicIntegrity"
	AndroidManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType_BasicIntegrityAndDeviceCertification AndroidManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType = "basicIntegrityAndDeviceCertification"
	AndroidManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType_None                                 AndroidManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType = "none"
)

func PossibleValuesForAndroidManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType() []string {
	return []string{
		string(AndroidManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType_BasicIntegrity),
		string(AndroidManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType_BasicIntegrityAndDeviceCertification),
		string(AndroidManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType_None),
	}
}

func (s *AndroidManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType(input string) (*AndroidManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType, error) {
	vals := map[string]AndroidManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType{
		"basicintegrity":                       AndroidManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType_BasicIntegrity,
		"basicintegrityanddevicecertification": AndroidManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType_BasicIntegrityAndDeviceCertification,
		"none":                                 AndroidManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType(input)
	return &out, nil
}
