package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType string

const (
	DefaultManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType_BasicIntegrity                       DefaultManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType = "basicIntegrity"
	DefaultManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType_BasicIntegrityAndDeviceCertification DefaultManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType = "basicIntegrityAndDeviceCertification"
	DefaultManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType_None                                 DefaultManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType = "none"
)

func PossibleValuesForDefaultManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType() []string {
	return []string{
		string(DefaultManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType_BasicIntegrity),
		string(DefaultManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType_BasicIntegrityAndDeviceCertification),
		string(DefaultManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType_None),
	}
}

func (s *DefaultManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDefaultManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDefaultManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType(input string) (*DefaultManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType, error) {
	vals := map[string]DefaultManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType{
		"basicintegrity":                       DefaultManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType_BasicIntegrity,
		"basicintegrityanddevicecertification": DefaultManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType_BasicIntegrityAndDeviceCertification,
		"none":                                 DefaultManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType(input)
	return &out, nil
}
