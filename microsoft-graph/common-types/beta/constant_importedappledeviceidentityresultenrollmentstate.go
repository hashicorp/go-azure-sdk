package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImportedAppleDeviceIdentityResultEnrollmentState string

const (
	ImportedAppleDeviceIdentityResultEnrollmentState_Blocked      ImportedAppleDeviceIdentityResultEnrollmentState = "blocked"
	ImportedAppleDeviceIdentityResultEnrollmentState_Enrolled     ImportedAppleDeviceIdentityResultEnrollmentState = "enrolled"
	ImportedAppleDeviceIdentityResultEnrollmentState_Failed       ImportedAppleDeviceIdentityResultEnrollmentState = "failed"
	ImportedAppleDeviceIdentityResultEnrollmentState_NotContacted ImportedAppleDeviceIdentityResultEnrollmentState = "notContacted"
	ImportedAppleDeviceIdentityResultEnrollmentState_PendingReset ImportedAppleDeviceIdentityResultEnrollmentState = "pendingReset"
	ImportedAppleDeviceIdentityResultEnrollmentState_Unknown      ImportedAppleDeviceIdentityResultEnrollmentState = "unknown"
)

func PossibleValuesForImportedAppleDeviceIdentityResultEnrollmentState() []string {
	return []string{
		string(ImportedAppleDeviceIdentityResultEnrollmentState_Blocked),
		string(ImportedAppleDeviceIdentityResultEnrollmentState_Enrolled),
		string(ImportedAppleDeviceIdentityResultEnrollmentState_Failed),
		string(ImportedAppleDeviceIdentityResultEnrollmentState_NotContacted),
		string(ImportedAppleDeviceIdentityResultEnrollmentState_PendingReset),
		string(ImportedAppleDeviceIdentityResultEnrollmentState_Unknown),
	}
}

func (s *ImportedAppleDeviceIdentityResultEnrollmentState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseImportedAppleDeviceIdentityResultEnrollmentState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseImportedAppleDeviceIdentityResultEnrollmentState(input string) (*ImportedAppleDeviceIdentityResultEnrollmentState, error) {
	vals := map[string]ImportedAppleDeviceIdentityResultEnrollmentState{
		"blocked":      ImportedAppleDeviceIdentityResultEnrollmentState_Blocked,
		"enrolled":     ImportedAppleDeviceIdentityResultEnrollmentState_Enrolled,
		"failed":       ImportedAppleDeviceIdentityResultEnrollmentState_Failed,
		"notcontacted": ImportedAppleDeviceIdentityResultEnrollmentState_NotContacted,
		"pendingreset": ImportedAppleDeviceIdentityResultEnrollmentState_PendingReset,
		"unknown":      ImportedAppleDeviceIdentityResultEnrollmentState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ImportedAppleDeviceIdentityResultEnrollmentState(input)
	return &out, nil
}
