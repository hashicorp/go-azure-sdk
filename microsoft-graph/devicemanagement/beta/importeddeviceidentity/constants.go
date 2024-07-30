package importeddeviceidentity

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImportedDeviceIdentityEnrollmentState string

const (
	ImportedDeviceIdentityEnrollmentStateBlocked      ImportedDeviceIdentityEnrollmentState = "blocked"
	ImportedDeviceIdentityEnrollmentStateEnrolled     ImportedDeviceIdentityEnrollmentState = "enrolled"
	ImportedDeviceIdentityEnrollmentStateFailed       ImportedDeviceIdentityEnrollmentState = "failed"
	ImportedDeviceIdentityEnrollmentStateNotContacted ImportedDeviceIdentityEnrollmentState = "notContacted"
	ImportedDeviceIdentityEnrollmentStatePendingReset ImportedDeviceIdentityEnrollmentState = "pendingReset"
	ImportedDeviceIdentityEnrollmentStateUnknown      ImportedDeviceIdentityEnrollmentState = "unknown"
)

func PossibleValuesForImportedDeviceIdentityEnrollmentState() []string {
	return []string{
		string(ImportedDeviceIdentityEnrollmentStateBlocked),
		string(ImportedDeviceIdentityEnrollmentStateEnrolled),
		string(ImportedDeviceIdentityEnrollmentStateFailed),
		string(ImportedDeviceIdentityEnrollmentStateNotContacted),
		string(ImportedDeviceIdentityEnrollmentStatePendingReset),
		string(ImportedDeviceIdentityEnrollmentStateUnknown),
	}
}

func (s *ImportedDeviceIdentityEnrollmentState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseImportedDeviceIdentityEnrollmentState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseImportedDeviceIdentityEnrollmentState(input string) (*ImportedDeviceIdentityEnrollmentState, error) {
	vals := map[string]ImportedDeviceIdentityEnrollmentState{
		"blocked":      ImportedDeviceIdentityEnrollmentStateBlocked,
		"enrolled":     ImportedDeviceIdentityEnrollmentStateEnrolled,
		"failed":       ImportedDeviceIdentityEnrollmentStateFailed,
		"notcontacted": ImportedDeviceIdentityEnrollmentStateNotContacted,
		"pendingreset": ImportedDeviceIdentityEnrollmentStatePendingReset,
		"unknown":      ImportedDeviceIdentityEnrollmentStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ImportedDeviceIdentityEnrollmentState(input)
	return &out, nil
}

type ImportedDeviceIdentityImportedDeviceIdentityType string

const (
	ImportedDeviceIdentityImportedDeviceIdentityTypeImei         ImportedDeviceIdentityImportedDeviceIdentityType = "imei"
	ImportedDeviceIdentityImportedDeviceIdentityTypeSerialNumber ImportedDeviceIdentityImportedDeviceIdentityType = "serialNumber"
	ImportedDeviceIdentityImportedDeviceIdentityTypeUnknown      ImportedDeviceIdentityImportedDeviceIdentityType = "unknown"
)

func PossibleValuesForImportedDeviceIdentityImportedDeviceIdentityType() []string {
	return []string{
		string(ImportedDeviceIdentityImportedDeviceIdentityTypeImei),
		string(ImportedDeviceIdentityImportedDeviceIdentityTypeSerialNumber),
		string(ImportedDeviceIdentityImportedDeviceIdentityTypeUnknown),
	}
}

func (s *ImportedDeviceIdentityImportedDeviceIdentityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseImportedDeviceIdentityImportedDeviceIdentityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseImportedDeviceIdentityImportedDeviceIdentityType(input string) (*ImportedDeviceIdentityImportedDeviceIdentityType, error) {
	vals := map[string]ImportedDeviceIdentityImportedDeviceIdentityType{
		"imei":         ImportedDeviceIdentityImportedDeviceIdentityTypeImei,
		"serialnumber": ImportedDeviceIdentityImportedDeviceIdentityTypeSerialNumber,
		"unknown":      ImportedDeviceIdentityImportedDeviceIdentityTypeUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ImportedDeviceIdentityImportedDeviceIdentityType(input)
	return &out, nil
}

type ImportedDeviceIdentityPlatform string

const (
	ImportedDeviceIdentityPlatformAndroid       ImportedDeviceIdentityPlatform = "android"
	ImportedDeviceIdentityPlatformIos           ImportedDeviceIdentityPlatform = "ios"
	ImportedDeviceIdentityPlatformMacOS         ImportedDeviceIdentityPlatform = "macOS"
	ImportedDeviceIdentityPlatformUnknown       ImportedDeviceIdentityPlatform = "unknown"
	ImportedDeviceIdentityPlatformWindows       ImportedDeviceIdentityPlatform = "windows"
	ImportedDeviceIdentityPlatformWindowsMobile ImportedDeviceIdentityPlatform = "windowsMobile"
)

func PossibleValuesForImportedDeviceIdentityPlatform() []string {
	return []string{
		string(ImportedDeviceIdentityPlatformAndroid),
		string(ImportedDeviceIdentityPlatformIos),
		string(ImportedDeviceIdentityPlatformMacOS),
		string(ImportedDeviceIdentityPlatformUnknown),
		string(ImportedDeviceIdentityPlatformWindows),
		string(ImportedDeviceIdentityPlatformWindowsMobile),
	}
}

func (s *ImportedDeviceIdentityPlatform) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseImportedDeviceIdentityPlatform(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseImportedDeviceIdentityPlatform(input string) (*ImportedDeviceIdentityPlatform, error) {
	vals := map[string]ImportedDeviceIdentityPlatform{
		"android":       ImportedDeviceIdentityPlatformAndroid,
		"ios":           ImportedDeviceIdentityPlatformIos,
		"macos":         ImportedDeviceIdentityPlatformMacOS,
		"unknown":       ImportedDeviceIdentityPlatformUnknown,
		"windows":       ImportedDeviceIdentityPlatformWindows,
		"windowsmobile": ImportedDeviceIdentityPlatformWindowsMobile,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ImportedDeviceIdentityPlatform(input)
	return &out, nil
}
