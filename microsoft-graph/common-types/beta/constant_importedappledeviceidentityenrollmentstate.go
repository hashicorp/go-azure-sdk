package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImportedAppleDeviceIdentityEnrollmentState string

const (
	ImportedAppleDeviceIdentityEnrollmentState_Blocked      ImportedAppleDeviceIdentityEnrollmentState = "blocked"
	ImportedAppleDeviceIdentityEnrollmentState_Enrolled     ImportedAppleDeviceIdentityEnrollmentState = "enrolled"
	ImportedAppleDeviceIdentityEnrollmentState_Failed       ImportedAppleDeviceIdentityEnrollmentState = "failed"
	ImportedAppleDeviceIdentityEnrollmentState_NotContacted ImportedAppleDeviceIdentityEnrollmentState = "notContacted"
	ImportedAppleDeviceIdentityEnrollmentState_PendingReset ImportedAppleDeviceIdentityEnrollmentState = "pendingReset"
	ImportedAppleDeviceIdentityEnrollmentState_Unknown      ImportedAppleDeviceIdentityEnrollmentState = "unknown"
)

func PossibleValuesForImportedAppleDeviceIdentityEnrollmentState() []string {
	return []string{
		string(ImportedAppleDeviceIdentityEnrollmentState_Blocked),
		string(ImportedAppleDeviceIdentityEnrollmentState_Enrolled),
		string(ImportedAppleDeviceIdentityEnrollmentState_Failed),
		string(ImportedAppleDeviceIdentityEnrollmentState_NotContacted),
		string(ImportedAppleDeviceIdentityEnrollmentState_PendingReset),
		string(ImportedAppleDeviceIdentityEnrollmentState_Unknown),
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
		"blocked":      ImportedAppleDeviceIdentityEnrollmentState_Blocked,
		"enrolled":     ImportedAppleDeviceIdentityEnrollmentState_Enrolled,
		"failed":       ImportedAppleDeviceIdentityEnrollmentState_Failed,
		"notcontacted": ImportedAppleDeviceIdentityEnrollmentState_NotContacted,
		"pendingreset": ImportedAppleDeviceIdentityEnrollmentState_PendingReset,
		"unknown":      ImportedAppleDeviceIdentityEnrollmentState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ImportedAppleDeviceIdentityEnrollmentState(input)
	return &out, nil
}
