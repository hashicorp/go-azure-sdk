package policy

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttestationType string

const (
	AttestationTypeOpenEnclave AttestationType = "OpenEnclave"
	AttestationTypeSevSnpVM    AttestationType = "SevSnpVm"
	AttestationTypeSgxEnclave  AttestationType = "SgxEnclave"
	AttestationTypeTpm         AttestationType = "Tpm"
)

func PossibleValuesForAttestationType() []string {
	return []string{
		string(AttestationTypeOpenEnclave),
		string(AttestationTypeSevSnpVM),
		string(AttestationTypeSgxEnclave),
		string(AttestationTypeTpm),
	}
}

func (s *AttestationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAttestationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAttestationType(input string) (*AttestationType, error) {
	vals := map[string]AttestationType{
		"openenclave": AttestationTypeOpenEnclave,
		"sevsnpvm":    AttestationTypeSevSnpVM,
		"sgxenclave":  AttestationTypeSgxEnclave,
		"tpm":         AttestationTypeTpm,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AttestationType(input)
	return &out, nil
}
