package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DepEnrollmentProfileITunesPairingMode string

const (
	DepEnrollmentProfileITunesPairingMode_Allow               DepEnrollmentProfileITunesPairingMode = "allow"
	DepEnrollmentProfileITunesPairingMode_Disallow            DepEnrollmentProfileITunesPairingMode = "disallow"
	DepEnrollmentProfileITunesPairingMode_RequiresCertificate DepEnrollmentProfileITunesPairingMode = "requiresCertificate"
)

func PossibleValuesForDepEnrollmentProfileITunesPairingMode() []string {
	return []string{
		string(DepEnrollmentProfileITunesPairingMode_Allow),
		string(DepEnrollmentProfileITunesPairingMode_Disallow),
		string(DepEnrollmentProfileITunesPairingMode_RequiresCertificate),
	}
}

func (s *DepEnrollmentProfileITunesPairingMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDepEnrollmentProfileITunesPairingMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDepEnrollmentProfileITunesPairingMode(input string) (*DepEnrollmentProfileITunesPairingMode, error) {
	vals := map[string]DepEnrollmentProfileITunesPairingMode{
		"allow":               DepEnrollmentProfileITunesPairingMode_Allow,
		"disallow":            DepEnrollmentProfileITunesPairingMode_Disallow,
		"requirescertificate": DepEnrollmentProfileITunesPairingMode_RequiresCertificate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DepEnrollmentProfileITunesPairingMode(input)
	return &out, nil
}
