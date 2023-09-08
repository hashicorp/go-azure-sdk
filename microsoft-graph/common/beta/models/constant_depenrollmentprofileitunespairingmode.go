package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DepEnrollmentProfileITunesPairingMode string

const (
	DepEnrollmentProfileITunesPairingModeallow               DepEnrollmentProfileITunesPairingMode = "Allow"
	DepEnrollmentProfileITunesPairingModedisallow            DepEnrollmentProfileITunesPairingMode = "Disallow"
	DepEnrollmentProfileITunesPairingModerequiresCertificate DepEnrollmentProfileITunesPairingMode = "RequiresCertificate"
)

func PossibleValuesForDepEnrollmentProfileITunesPairingMode() []string {
	return []string{
		string(DepEnrollmentProfileITunesPairingModeallow),
		string(DepEnrollmentProfileITunesPairingModedisallow),
		string(DepEnrollmentProfileITunesPairingModerequiresCertificate),
	}
}

func parseDepEnrollmentProfileITunesPairingMode(input string) (*DepEnrollmentProfileITunesPairingMode, error) {
	vals := map[string]DepEnrollmentProfileITunesPairingMode{
		"allow":               DepEnrollmentProfileITunesPairingModeallow,
		"disallow":            DepEnrollmentProfileITunesPairingModedisallow,
		"requirescertificate": DepEnrollmentProfileITunesPairingModerequiresCertificate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DepEnrollmentProfileITunesPairingMode(input)
	return &out, nil
}
