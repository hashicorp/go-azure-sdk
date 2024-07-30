package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImportedDeviceIdentityResultEnrollmentState string

const (
	ImportedDeviceIdentityResultEnrollmentState_Blocked      ImportedDeviceIdentityResultEnrollmentState = "blocked"
	ImportedDeviceIdentityResultEnrollmentState_Enrolled     ImportedDeviceIdentityResultEnrollmentState = "enrolled"
	ImportedDeviceIdentityResultEnrollmentState_Failed       ImportedDeviceIdentityResultEnrollmentState = "failed"
	ImportedDeviceIdentityResultEnrollmentState_NotContacted ImportedDeviceIdentityResultEnrollmentState = "notContacted"
	ImportedDeviceIdentityResultEnrollmentState_PendingReset ImportedDeviceIdentityResultEnrollmentState = "pendingReset"
	ImportedDeviceIdentityResultEnrollmentState_Unknown      ImportedDeviceIdentityResultEnrollmentState = "unknown"
)

func PossibleValuesForImportedDeviceIdentityResultEnrollmentState() []string {
	return []string{
		string(ImportedDeviceIdentityResultEnrollmentState_Blocked),
		string(ImportedDeviceIdentityResultEnrollmentState_Enrolled),
		string(ImportedDeviceIdentityResultEnrollmentState_Failed),
		string(ImportedDeviceIdentityResultEnrollmentState_NotContacted),
		string(ImportedDeviceIdentityResultEnrollmentState_PendingReset),
		string(ImportedDeviceIdentityResultEnrollmentState_Unknown),
	}
}

func (s *ImportedDeviceIdentityResultEnrollmentState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseImportedDeviceIdentityResultEnrollmentState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseImportedDeviceIdentityResultEnrollmentState(input string) (*ImportedDeviceIdentityResultEnrollmentState, error) {
	vals := map[string]ImportedDeviceIdentityResultEnrollmentState{
		"blocked":      ImportedDeviceIdentityResultEnrollmentState_Blocked,
		"enrolled":     ImportedDeviceIdentityResultEnrollmentState_Enrolled,
		"failed":       ImportedDeviceIdentityResultEnrollmentState_Failed,
		"notcontacted": ImportedDeviceIdentityResultEnrollmentState_NotContacted,
		"pendingreset": ImportedDeviceIdentityResultEnrollmentState_PendingReset,
		"unknown":      ImportedDeviceIdentityResultEnrollmentState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ImportedDeviceIdentityResultEnrollmentState(input)
	return &out, nil
}
