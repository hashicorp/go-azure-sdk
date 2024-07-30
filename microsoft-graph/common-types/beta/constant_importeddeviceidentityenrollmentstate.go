package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImportedDeviceIdentityEnrollmentState string

const (
	ImportedDeviceIdentityEnrollmentState_Blocked      ImportedDeviceIdentityEnrollmentState = "blocked"
	ImportedDeviceIdentityEnrollmentState_Enrolled     ImportedDeviceIdentityEnrollmentState = "enrolled"
	ImportedDeviceIdentityEnrollmentState_Failed       ImportedDeviceIdentityEnrollmentState = "failed"
	ImportedDeviceIdentityEnrollmentState_NotContacted ImportedDeviceIdentityEnrollmentState = "notContacted"
	ImportedDeviceIdentityEnrollmentState_PendingReset ImportedDeviceIdentityEnrollmentState = "pendingReset"
	ImportedDeviceIdentityEnrollmentState_Unknown      ImportedDeviceIdentityEnrollmentState = "unknown"
)

func PossibleValuesForImportedDeviceIdentityEnrollmentState() []string {
	return []string{
		string(ImportedDeviceIdentityEnrollmentState_Blocked),
		string(ImportedDeviceIdentityEnrollmentState_Enrolled),
		string(ImportedDeviceIdentityEnrollmentState_Failed),
		string(ImportedDeviceIdentityEnrollmentState_NotContacted),
		string(ImportedDeviceIdentityEnrollmentState_PendingReset),
		string(ImportedDeviceIdentityEnrollmentState_Unknown),
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
		"blocked":      ImportedDeviceIdentityEnrollmentState_Blocked,
		"enrolled":     ImportedDeviceIdentityEnrollmentState_Enrolled,
		"failed":       ImportedDeviceIdentityEnrollmentState_Failed,
		"notcontacted": ImportedDeviceIdentityEnrollmentState_NotContacted,
		"pendingreset": ImportedDeviceIdentityEnrollmentState_PendingReset,
		"unknown":      ImportedDeviceIdentityEnrollmentState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ImportedDeviceIdentityEnrollmentState(input)
	return &out, nil
}
