package deponboardingsettingimportedappledeviceidentity

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImportedAppleDeviceIdentityDiscoverySource string

const (
	ImportedAppleDeviceIdentityDiscoverySourceAdminImport             ImportedAppleDeviceIdentityDiscoverySource = "adminImport"
	ImportedAppleDeviceIdentityDiscoverySourceDeviceEnrollmentProgram ImportedAppleDeviceIdentityDiscoverySource = "deviceEnrollmentProgram"
	ImportedAppleDeviceIdentityDiscoverySourceUnknown                 ImportedAppleDeviceIdentityDiscoverySource = "unknown"
)

func PossibleValuesForImportedAppleDeviceIdentityDiscoverySource() []string {
	return []string{
		string(ImportedAppleDeviceIdentityDiscoverySourceAdminImport),
		string(ImportedAppleDeviceIdentityDiscoverySourceDeviceEnrollmentProgram),
		string(ImportedAppleDeviceIdentityDiscoverySourceUnknown),
	}
}

func (s *ImportedAppleDeviceIdentityDiscoverySource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseImportedAppleDeviceIdentityDiscoverySource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseImportedAppleDeviceIdentityDiscoverySource(input string) (*ImportedAppleDeviceIdentityDiscoverySource, error) {
	vals := map[string]ImportedAppleDeviceIdentityDiscoverySource{
		"adminimport":             ImportedAppleDeviceIdentityDiscoverySourceAdminImport,
		"deviceenrollmentprogram": ImportedAppleDeviceIdentityDiscoverySourceDeviceEnrollmentProgram,
		"unknown":                 ImportedAppleDeviceIdentityDiscoverySourceUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ImportedAppleDeviceIdentityDiscoverySource(input)
	return &out, nil
}

type ImportedAppleDeviceIdentityEnrollmentState string

const (
	ImportedAppleDeviceIdentityEnrollmentStateBlocked      ImportedAppleDeviceIdentityEnrollmentState = "blocked"
	ImportedAppleDeviceIdentityEnrollmentStateEnrolled     ImportedAppleDeviceIdentityEnrollmentState = "enrolled"
	ImportedAppleDeviceIdentityEnrollmentStateFailed       ImportedAppleDeviceIdentityEnrollmentState = "failed"
	ImportedAppleDeviceIdentityEnrollmentStateNotContacted ImportedAppleDeviceIdentityEnrollmentState = "notContacted"
	ImportedAppleDeviceIdentityEnrollmentStatePendingReset ImportedAppleDeviceIdentityEnrollmentState = "pendingReset"
	ImportedAppleDeviceIdentityEnrollmentStateUnknown      ImportedAppleDeviceIdentityEnrollmentState = "unknown"
)

func PossibleValuesForImportedAppleDeviceIdentityEnrollmentState() []string {
	return []string{
		string(ImportedAppleDeviceIdentityEnrollmentStateBlocked),
		string(ImportedAppleDeviceIdentityEnrollmentStateEnrolled),
		string(ImportedAppleDeviceIdentityEnrollmentStateFailed),
		string(ImportedAppleDeviceIdentityEnrollmentStateNotContacted),
		string(ImportedAppleDeviceIdentityEnrollmentStatePendingReset),
		string(ImportedAppleDeviceIdentityEnrollmentStateUnknown),
	}
}

func (s *ImportedAppleDeviceIdentityEnrollmentState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseImportedAppleDeviceIdentityEnrollmentState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseImportedAppleDeviceIdentityEnrollmentState(input string) (*ImportedAppleDeviceIdentityEnrollmentState, error) {
	vals := map[string]ImportedAppleDeviceIdentityEnrollmentState{
		"blocked":      ImportedAppleDeviceIdentityEnrollmentStateBlocked,
		"enrolled":     ImportedAppleDeviceIdentityEnrollmentStateEnrolled,
		"failed":       ImportedAppleDeviceIdentityEnrollmentStateFailed,
		"notcontacted": ImportedAppleDeviceIdentityEnrollmentStateNotContacted,
		"pendingreset": ImportedAppleDeviceIdentityEnrollmentStatePendingReset,
		"unknown":      ImportedAppleDeviceIdentityEnrollmentStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ImportedAppleDeviceIdentityEnrollmentState(input)
	return &out, nil
}

type ImportedAppleDeviceIdentityPlatform string

const (
	ImportedAppleDeviceIdentityPlatformAndroid       ImportedAppleDeviceIdentityPlatform = "android"
	ImportedAppleDeviceIdentityPlatformIos           ImportedAppleDeviceIdentityPlatform = "ios"
	ImportedAppleDeviceIdentityPlatformMacOS         ImportedAppleDeviceIdentityPlatform = "macOS"
	ImportedAppleDeviceIdentityPlatformUnknown       ImportedAppleDeviceIdentityPlatform = "unknown"
	ImportedAppleDeviceIdentityPlatformWindows       ImportedAppleDeviceIdentityPlatform = "windows"
	ImportedAppleDeviceIdentityPlatformWindowsMobile ImportedAppleDeviceIdentityPlatform = "windowsMobile"
)

func PossibleValuesForImportedAppleDeviceIdentityPlatform() []string {
	return []string{
		string(ImportedAppleDeviceIdentityPlatformAndroid),
		string(ImportedAppleDeviceIdentityPlatformIos),
		string(ImportedAppleDeviceIdentityPlatformMacOS),
		string(ImportedAppleDeviceIdentityPlatformUnknown),
		string(ImportedAppleDeviceIdentityPlatformWindows),
		string(ImportedAppleDeviceIdentityPlatformWindowsMobile),
	}
}

func (s *ImportedAppleDeviceIdentityPlatform) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseImportedAppleDeviceIdentityPlatform(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseImportedAppleDeviceIdentityPlatform(input string) (*ImportedAppleDeviceIdentityPlatform, error) {
	vals := map[string]ImportedAppleDeviceIdentityPlatform{
		"android":       ImportedAppleDeviceIdentityPlatformAndroid,
		"ios":           ImportedAppleDeviceIdentityPlatformIos,
		"macos":         ImportedAppleDeviceIdentityPlatformMacOS,
		"unknown":       ImportedAppleDeviceIdentityPlatformUnknown,
		"windows":       ImportedAppleDeviceIdentityPlatformWindows,
		"windowsmobile": ImportedAppleDeviceIdentityPlatformWindowsMobile,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ImportedAppleDeviceIdentityPlatform(input)
	return &out, nil
}
