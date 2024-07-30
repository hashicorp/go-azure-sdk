package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DepIOSEnrollmentProfileITunesPairingMode string

const (
	DepIOSEnrollmentProfileITunesPairingMode_Allow               DepIOSEnrollmentProfileITunesPairingMode = "allow"
	DepIOSEnrollmentProfileITunesPairingMode_Disallow            DepIOSEnrollmentProfileITunesPairingMode = "disallow"
	DepIOSEnrollmentProfileITunesPairingMode_RequiresCertificate DepIOSEnrollmentProfileITunesPairingMode = "requiresCertificate"
)

func PossibleValuesForDepIOSEnrollmentProfileITunesPairingMode() []string {
	return []string{
		string(DepIOSEnrollmentProfileITunesPairingMode_Allow),
		string(DepIOSEnrollmentProfileITunesPairingMode_Disallow),
		string(DepIOSEnrollmentProfileITunesPairingMode_RequiresCertificate),
	}
}

func (s *DepIOSEnrollmentProfileITunesPairingMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDepIOSEnrollmentProfileITunesPairingMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDepIOSEnrollmentProfileITunesPairingMode(input string) (*DepIOSEnrollmentProfileITunesPairingMode, error) {
	vals := map[string]DepIOSEnrollmentProfileITunesPairingMode{
		"allow":               DepIOSEnrollmentProfileITunesPairingMode_Allow,
		"disallow":            DepIOSEnrollmentProfileITunesPairingMode_Disallow,
		"requirescertificate": DepIOSEnrollmentProfileITunesPairingMode_RequiresCertificate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DepIOSEnrollmentProfileITunesPairingMode(input)
	return &out, nil
}
