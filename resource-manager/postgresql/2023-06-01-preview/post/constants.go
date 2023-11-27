package post

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MigrationNameAvailabilityReason string

const (
	MigrationNameAvailabilityReasonAlreadyExists MigrationNameAvailabilityReason = "AlreadyExists"
	MigrationNameAvailabilityReasonInvalid       MigrationNameAvailabilityReason = "Invalid"
)

func PossibleValuesForMigrationNameAvailabilityReason() []string {
	return []string{
		string(MigrationNameAvailabilityReasonAlreadyExists),
		string(MigrationNameAvailabilityReasonInvalid),
	}
}

func (s *MigrationNameAvailabilityReason) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMigrationNameAvailabilityReason(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMigrationNameAvailabilityReason(input string) (*MigrationNameAvailabilityReason, error) {
	vals := map[string]MigrationNameAvailabilityReason{
		"alreadyexists": MigrationNameAvailabilityReasonAlreadyExists,
		"invalid":       MigrationNameAvailabilityReasonInvalid,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MigrationNameAvailabilityReason(input)
	return &out, nil
}
