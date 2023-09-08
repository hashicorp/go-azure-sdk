package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Fido2AuthenticationMethodAttestationLevel string

const (
	Fido2AuthenticationMethodAttestationLevelattested    Fido2AuthenticationMethodAttestationLevel = "Attested"
	Fido2AuthenticationMethodAttestationLevelnotAttested Fido2AuthenticationMethodAttestationLevel = "NotAttested"
)

func PossibleValuesForFido2AuthenticationMethodAttestationLevel() []string {
	return []string{
		string(Fido2AuthenticationMethodAttestationLevelattested),
		string(Fido2AuthenticationMethodAttestationLevelnotAttested),
	}
}

func parseFido2AuthenticationMethodAttestationLevel(input string) (*Fido2AuthenticationMethodAttestationLevel, error) {
	vals := map[string]Fido2AuthenticationMethodAttestationLevel{
		"attested":    Fido2AuthenticationMethodAttestationLevelattested,
		"notattested": Fido2AuthenticationMethodAttestationLevelnotAttested,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Fido2AuthenticationMethodAttestationLevel(input)
	return &out, nil
}
