package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DepIOSEnrollmentProfileITunesPairingMode string

const (
	DepIOSEnrollmentProfileITunesPairingModeallow               DepIOSEnrollmentProfileITunesPairingMode = "Allow"
	DepIOSEnrollmentProfileITunesPairingModedisallow            DepIOSEnrollmentProfileITunesPairingMode = "Disallow"
	DepIOSEnrollmentProfileITunesPairingModerequiresCertificate DepIOSEnrollmentProfileITunesPairingMode = "RequiresCertificate"
)

func PossibleValuesForDepIOSEnrollmentProfileITunesPairingMode() []string {
	return []string{
		string(DepIOSEnrollmentProfileITunesPairingModeallow),
		string(DepIOSEnrollmentProfileITunesPairingModedisallow),
		string(DepIOSEnrollmentProfileITunesPairingModerequiresCertificate),
	}
}

func parseDepIOSEnrollmentProfileITunesPairingMode(input string) (*DepIOSEnrollmentProfileITunesPairingMode, error) {
	vals := map[string]DepIOSEnrollmentProfileITunesPairingMode{
		"allow":               DepIOSEnrollmentProfileITunesPairingModeallow,
		"disallow":            DepIOSEnrollmentProfileITunesPairingModedisallow,
		"requirescertificate": DepIOSEnrollmentProfileITunesPairingModerequiresCertificate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DepIOSEnrollmentProfileITunesPairingMode(input)
	return &out, nil
}
