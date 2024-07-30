package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Fido2AuthenticationMethodAttestationLevel string

const (
	Fido2AuthenticationMethodAttestationLevel_Attested    Fido2AuthenticationMethodAttestationLevel = "attested"
	Fido2AuthenticationMethodAttestationLevel_NotAttested Fido2AuthenticationMethodAttestationLevel = "notAttested"
)

func PossibleValuesForFido2AuthenticationMethodAttestationLevel() []string {
	return []string{
		string(Fido2AuthenticationMethodAttestationLevel_Attested),
		string(Fido2AuthenticationMethodAttestationLevel_NotAttested),
	}
}

func (s *Fido2AuthenticationMethodAttestationLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFido2AuthenticationMethodAttestationLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFido2AuthenticationMethodAttestationLevel(input string) (*Fido2AuthenticationMethodAttestationLevel, error) {
	vals := map[string]Fido2AuthenticationMethodAttestationLevel{
		"attested":    Fido2AuthenticationMethodAttestationLevel_Attested,
		"notattested": Fido2AuthenticationMethodAttestationLevel_NotAttested,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Fido2AuthenticationMethodAttestationLevel(input)
	return &out, nil
}
