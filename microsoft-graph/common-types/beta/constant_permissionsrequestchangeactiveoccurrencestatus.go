package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PermissionsRequestChangeActiveOccurrenceStatus string

const (
	PermissionsRequestChangeActiveOccurrenceStatus_Granted        PermissionsRequestChangeActiveOccurrenceStatus = "granted"
	PermissionsRequestChangeActiveOccurrenceStatus_Granting       PermissionsRequestChangeActiveOccurrenceStatus = "granting"
	PermissionsRequestChangeActiveOccurrenceStatus_GrantingFailed PermissionsRequestChangeActiveOccurrenceStatus = "grantingFailed"
	PermissionsRequestChangeActiveOccurrenceStatus_Revoked        PermissionsRequestChangeActiveOccurrenceStatus = "revoked"
	PermissionsRequestChangeActiveOccurrenceStatus_Revoking       PermissionsRequestChangeActiveOccurrenceStatus = "revoking"
	PermissionsRequestChangeActiveOccurrenceStatus_RevokingFailed PermissionsRequestChangeActiveOccurrenceStatus = "revokingFailed"
)

func PossibleValuesForPermissionsRequestChangeActiveOccurrenceStatus() []string {
	return []string{
		string(PermissionsRequestChangeActiveOccurrenceStatus_Granted),
		string(PermissionsRequestChangeActiveOccurrenceStatus_Granting),
		string(PermissionsRequestChangeActiveOccurrenceStatus_GrantingFailed),
		string(PermissionsRequestChangeActiveOccurrenceStatus_Revoked),
		string(PermissionsRequestChangeActiveOccurrenceStatus_Revoking),
		string(PermissionsRequestChangeActiveOccurrenceStatus_RevokingFailed),
	}
}

func (s *PermissionsRequestChangeActiveOccurrenceStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePermissionsRequestChangeActiveOccurrenceStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePermissionsRequestChangeActiveOccurrenceStatus(input string) (*PermissionsRequestChangeActiveOccurrenceStatus, error) {
	vals := map[string]PermissionsRequestChangeActiveOccurrenceStatus{
		"granted":        PermissionsRequestChangeActiveOccurrenceStatus_Granted,
		"granting":       PermissionsRequestChangeActiveOccurrenceStatus_Granting,
		"grantingfailed": PermissionsRequestChangeActiveOccurrenceStatus_GrantingFailed,
		"revoked":        PermissionsRequestChangeActiveOccurrenceStatus_Revoked,
		"revoking":       PermissionsRequestChangeActiveOccurrenceStatus_Revoking,
		"revokingfailed": PermissionsRequestChangeActiveOccurrenceStatus_RevokingFailed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PermissionsRequestChangeActiveOccurrenceStatus(input)
	return &out, nil
}
